package main

func main() {
	tasks := []Task{
		&EmailTask{Email: "email1@mail.com"},
		&ImageTask{ImageUrl: "/images/img1.jpg"},
		&EmailTask{Email: "email2@mail.com"},
		&ImageTask{ImageUrl: "/images/img2.jpg"},
		&EmailTask{Email: "email3@mail.com"},
		&ImageTask{ImageUrl: "/images/img3.jpg"},
		&EmailTask{Email: "email4@mail.com"},
		&ImageTask{ImageUrl: "/images/img4.jpg"},
		&EmailTask{Email: "email5@mail.com"},
		&ImageTask{ImageUrl: "/images/img5.jpg"},
		&ImageTask{ImageUrl: "/images/img6.jpg"},
	}

	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()
}
