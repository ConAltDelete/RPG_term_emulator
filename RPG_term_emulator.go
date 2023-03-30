package RPG_term_emulator

type Place struct {
	Place_id int
	// Contain pointer(s) to files with relevant info
}

type Context struct {
	context_code int
	Place_info Place
}

func (C Context) login_check(user string, password string){
	print("Granted")
}

func login_check(user string, password string, place string) (err int, code int){
	return 0,0
}

func main(){
	for;;{ // Read Evaluate Execute Loop
		
	}
}