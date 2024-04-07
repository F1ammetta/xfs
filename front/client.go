package front

import (
	"fmt"
	// "github.com/lxn/win"
	"github.com/webview/webview_go"
)

// func getScreenSize() (int, int) {
// 	return int(win.GetSystemMetrics(win.SM_CXSCREEN)), int(win.GetSystemMetrics(win.SM_CYSCREEN))
// }

func Run(port int) {
	// could be abstracted but only uses winapi for now
	// width, height := getScreenSize()
	//
	// width = int(float64(width) * 0.8)
	//
	// height = int(float64(height) * 0.8)

	w := webview.New(false)

	defer w.Destroy()

	// testing winapi stuff
	// hwnd := win.HWND(w.Window())
	//
	// style := win.GetWindowLong(hwnd, win.GWL_STYLE)
	// style &^= win.WS_CAPTION | win.WS_SYSMENU | win.WS_THICKFRAME | win.WS_MINIMIZEBOX | win.WS_MAXIMIZEBOX | win.PFM_BORDER
	// win.SetWindowLong(hwnd, win.GWL_STYLE, style)
	//
	// // with the title bar gone, we need to be able to drag the window around
	// win.SetWindowPos(hwnd, win.HWND_TOP, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE)

	w.SetTitle("XFS")

	// w.SetSize(width, height, webview.HintNone)

	w.Navigate(fmt.Sprintf("http://localhost:%d", port))

	w.Run()
}
