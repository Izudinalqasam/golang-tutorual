package basic

import "fmt"

// Defer function is a function we can schedule to be called after a function executed, if error happend, the defer function still will be executed
// Panic function is function which can stop program, but defer method will still be executed
// Recover is function which can catch data panic, with this, process panic will stop, so programm will continue to run
func MainDeferPanicRecover() {
	runApplication()
	runApp(true)
	runAppErrorCatch(true)
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func logging() {
	fmt.Println("Logging")
}

func endApp() {
	message := recover()
	fmt.Println("Error dengan message:", message)
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error happend")
	}

	fmt.Println("Aplikasi berjalan")
}

/**
 * @desc: runAppErrorCatch
 */
func runAppErrorCatch(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berjalan")
}
