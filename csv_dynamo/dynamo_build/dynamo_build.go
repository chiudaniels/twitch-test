// package dynamo_build
package build

import (
	"fmt"
    "os"
    "log"
	"bytes"
	"encoding/csv"
	"encoding/json"
	// "flag"
	// "io/ioutil"
	"strconv"
    "strings"
    "time"
    
	// "../convert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/service/athena"
	// "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	
)

type Compliance struct {
	Status string
    Installedtime string
    Executiontype string
    Patchseverity string
    Title string
    Severity string
    Executiontime string
    Compliancetype string
    Id string
    Documentversion string
    Patchstate string
    Patchbaselineid string
    Documentname string
    Patchgroup string
    Executionid string
    Resourceid string
    Capturetime string
    Schemaversion float32
    Accountid string
    Region string
    Resourcetype string
}

func getEntries() ([]Compliance, error){
    var entries []Compliance

    file_path, err := GetCSV()
    if err != nil{
        fmt.Println("Got error reading csv")
        fmt.Println(err.Error())
        return nil, err
    }

    json_entries := ReadCSV(file_path)
	json.Unmarshal(json_entries, &entries)
    // fmt.Println(entries[0])
    tableName := "compliance_items_test"

    sess, err := session.NewSession( &aws.Config{
        Region: aws.String("us-west-2"),
    })
    if err != nil{
        fmt.Println("got error making ession")
        fmt.Println(err.Error())
        // os.Exit(1)
        return nil, err
    }

    // sess := session.Must(session.NewSessionWithOptions(session.Options{
    //     SharedConfigState: session.SharedConfigEnable,
    // }))
    
    // Create DynamoDB client
    svc := dynamodb.New(sess)

    for _, item := range entries{
        av, err := dynamodbattribute.MarshalMap(item)
        if err != nil {
            fmt.Println("Got error marshalling map:")
            fmt.Println(err.Error())
            return nil, err
            // os.Exit(1)
        }
        input := &dynamodb.PutItemInput{
            Item:      av,
            TableName: aws.String(tableName),
        }

        _, err = svc.PutItem(input)
        if err != nil {
            fmt.Println("Got error calling PutItem:")
            fmt.Println(err.Error())
            return nil, err
            // os.Exit(1)
        }
    }
    
    fmt.Println("Added csv")

	return entries, nil
}

func GetCSV() (*os.File,  error) {
    bucket := "chiudani-pandora-test2"
    item_key := "accountid=307142784234/athena_result/118ac2b8-669e-42b7-9586-5f2d739c4791.csv"

    sess, err := session.NewSession( &aws.Config{
        Region: aws.String("us-west-2"),
    })
    if err != nil{
        fmt.Println("got error making session")
        fmt.Println(err.Error())
        // os.Exit(1)
        return nil, err
    }
    // 3) Create a new AWS S3 downloader 
    // buff := &aws.WriteAtBuffer{}
    
    downloader := s3manager.NewDownloader(sess)
    file, err := os.Create(item_key)

    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String(bucket),
            Key:    aws.String(item_key),
    })
    if err != nil {
        log.Fatalf("Unable to download item %q, %v", item_key, err)
    }

    fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

    // data := buff.Bytes()  


    return file, nil

}


func ReadCSV(path *os.File) ([]byte) {
	csvFile, err := os.Open(*path)

	if err != nil {
		log.Fatal("The file is not found || wrong root")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		log.Fatal("Something wrong, the file maybe empty or length of the lines are not the same")
	}

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	x, _ := json.MarshalIndent(rawMessage, "", "  ")
	return x
}

func AthenaGet(){
    sess, err := session.NewSession( &aws.Config{
        Region: aws.String("us-west-2"),
    })
    if err != nil{
        fmt.Println("got error making ession")
        fmt.Println(err.Error())
        // os.Exit(1)
        return nil, err
    }
    svc:= athena.New(sess)
    var s athena.StartQueryExecutionInput
    s.SetQueryString(`SELECT * FROM "crawler_test"."aws_complianceitem";`)

    var q athena.QueryExecutionContext
    q.SetDatabase("crawler_test")
    s.SetQueryExecutionContest(&q)

    var r athena.ResultConfiguration
    r.SetOutputLocation("s3://chiudani-pandora-test2/accountid=307142784234/athena_result/")
    result, err := svc.StartQueryExecution(&s)
	if err != nil {
		fmt.Println(err)
		return
    }
    
    fmt.Println("StartQueryExecution result:")
	fmt.Println(result.GoString())

	var qri athena.GetQueryExecutionInput
    qri.SetQueryExecutionId(*result.QueryExecutionId)
    
    var qrop *athena.GetQueryExecutionOutput
	duration := time.Duration(2) * time.Second // Pause for 2 secondsß

    for {
		qrop, err = svc.GetQueryExecution(&qri)
		if err != nil {
			fmt.Println(err)
			return
		}
		if *qrop.QueryExecution.Status.State != "RUNNING" {
			break
		}
		fmt.Println("waiting.")
		time.Sleep(duration)

	}
	if *qrop.QueryExecution.Status.State == "SUCCEEDED" {

		var ip athena.GetQueryResultsInput
		ip.SetQueryExecutionId(*result.QueryExecutionId)

		op, err := svc.GetQueryResults(&ip)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v", op)
	} else {
		fmt.Println(*qrop.QueryExecution.Status.State)

	}
}
// func main(){
// 	AthenaGet()
// }