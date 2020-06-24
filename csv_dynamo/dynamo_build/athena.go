package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

func AthenaGet(){
    sess, err := session.NewSession( &aws.Config{
        Region: aws.String("us-west-2"),
    })
    if err != nil{
        fmt.Println("got error making ession")
        fmt.Println(err.Error())
        // os.Exit(1)
        // return nil, err
    }
    svc:= athena.New(sess)
	var s athena.StartQueryExecutionInput
	
    s.SetQueryString(`SELECT * FROM "crawler_test"."aws_complianceitem" LIMIT 10;`)

	var q athena.QueryExecutionContext
	q.SetCatalog("AwsDataCatalog")
    // q.SetDatabase("crawler_test")
    s.SetQueryExecutionContext(&q)
	s.SetWorkGroup("primary")

    var r athena.ResultConfiguration
	r.SetOutputLocation("s3://chiudani-pandora-test2/accountid=307142784234/athena_result/")
	s.SetResultConfiguration(&r)

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
		fmt.Println("try executing run")
		qrop, err = svc.GetQueryExecution(&qri)
		if err != nil {
			fmt.Println(err)
			return
		}
		if *qrop.QueryExecution.Status.State != "RUNNING" && *qrop.QueryExecution.Status.State != "QUEUED"{
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
	} else if *qrop.QueryExecution.Status.State == "FAILED" {

		var failed_input athena.BatchGetQueryExecutionInput 
		var query []*string  
		query = append(query, result.QueryExecutionId)
		failed_input.SetQueryExecutionIds(query)

		// ip.SetQueryExecutionId(*result.QueryExecutionId)

		output, err := svc.BatchGetQueryExecution(&failed_input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Failed error:")
		fmt.Println(output.QueryExecutions[0].Status)
		// fmt.Printf("%+v", op)
	}  else {
		fmt.Println(*qrop.QueryExecution.Status.State)
	}
	
	
}


func AthenaGetDatabases(){
	sess, err := session.NewSession( &aws.Config{
        Region: aws.String("us-west-2"),
    })
    if err != nil{
        fmt.Println("got error making ession")
        fmt.Println(err.Error())
        // os.Exit(1)
        // return nil, err
    }
    svc:= athena.New(sess)
	var s athena.StartQueryExecutionInput
	
	var catalog_input athena.ListDataCatalogsInput
	catalogs , err := svc.ListDataCatalogs(&catalog_input)
	if err != nil{
        fmt.Println("error listing catalogs") 
        fmt.Println(err.Error())
        // os.Exit(1)
        // return nil, err
	}
	fmt.Println("listing catalogs")
	fmt.Println(catalogs.DataCatalogsSummary)
	
	var database_input athena.ListDatabasesInput
	// database_input.SetCatalogName("AwsDataCatalog")
	database_input.SetCatalogName(*catalogs.DataCatalogsSummary[0].CatalogName)
	fmt.Println(*catalogs.DataCatalogsSummary[0].CatalogName)
	databases, err := svc.ListDatabases(&database_input)
	if err != nil{
        fmt.Println("listing databases")
        fmt.Println(err.Error())
        // os.Exit(1)
        // return nil, err
	}
	fmt.Println("listing databases")
	fmt.Println(databases.DatabaseList)
}

func main(){
	AthenaGet()
}