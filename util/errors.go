package util

type Error struct {
	field string 
	message string
}

type Response struct {
	message string
	errors []Error
	data string
}


var (
	//BadRequestErr Response{message: "Bad request", }
)
