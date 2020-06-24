package servertest


import (
    "context"
    "fmt"
    // "github.com/twitchtv/twirp"
	pb "../rpc"
	"../model"
)

// Server implements the Haberdasher service
type Server struct {}

func (s *Server) GetAll(ctx context.Context, query *pb.GetAllQueryParam) (Movies *pb.Movies, err error) {
    entries := model.GetAllEntry()
	// json.NewEncoder(w).Encode(entries)
    NewMovies := pb.Movies{}
    for _,j := range entries{
        entry := pb.Movie{
            Title : j.Title,
            Year: int32(j.Year),
        }
        // fmt.Println(&entry)
        // fmt.Println(NewMovies.Data)
        NewMovies.Data = append(NewMovies.Data, &entry)
        // NewMovies.Data[i] = &entry
    }
    return &NewMovies, nil
}