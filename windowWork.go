package main

import (
	"embed"
	"fmt"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"image/png"
	"log"
)

var sampleData = map[string]string{
	"comp502": "Research\n(3 credits)\nPrerequisite: Consent of the department; formal application required\nOriginal research is undertaken by the graduate student in their field. This course culminates in a capstone project. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Offered fall and spring semesters.",
	"comp503": "Directed Study\n(1-3 credits)\nPrerequisite: Consent of the department; formal application required\nDirected study is designed for the graduate student who desires to study selected topics in a specific field. For details, consult the paragraph titled “Directed or Independent Study” in the “College of Graduate Studies” section of this catalog. Repeatable: may earn a maximum of six credits. Offered fall and spring semesters.",
	"comp510": "Topics in Programming Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course investigates programming language development from designer’s, user’s and implementer’s point of view. Topics include formal syntax and semantics, language system, extensible languages and control structures. There is also a survey of intralanguage features, covering ALGOL-60, ALGOL-68, Ada, Pascal, LISP, SNOBOL-4 APL, SIMULA-67, CLU, MODULA, and others. Offered periodically.",
	"comp520": "Operating Systems Principles\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course examines design principles such as optimal scheduling; file systems, system integrity and security, as well as the mathematical analysis of selected aspects of operating system design. Topics include queuing theory, disk scheduling, storage management and the working set model. Design and implementation of an operating system nucleus is also studied. Offered periodically.",
	"comp525": "Design and Construction of Compilers\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include lexical and syntactic analysis; code generation; error detection and correction; optimization techniques; models of code generators; and incremental and interactive compiling. Students will design and implement a compiler. Offered periodically.",
	"comp530": "Software Engineering\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include construction of reliable software, software tools, software testing methodologies, structured design, structured programming, software characteristics and quality and formal proofs of program correctness. Chief programmer teams and structure walk-throughs will be employed. Offered periodically.\n",
	"comp540": "Automata, Computability and Formal Languages\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nTopics in this course will include finite automata and regular languages, context- free languages, Turing machines and their variants, partial recursive functions and grammars, Church’s thesis, undecidable problems, complexity of algorithms and completeness. Offered periodically.",
	"comp545": "Analysis of Algorithms\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course deals with techniques in the analysis of algorithms. Topics to be chosen from among the following: dynamic programming, search and traverse techniques, backtracking, numerical techniques, NP-hard and NP-complete problems, approximation algorithms and other topics in the analysis and design of algorithms. Offered fall semester.\n",
	"comp560": "Artificial Intelligence\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to LISP or another AI programming language. Topics are chosen from pattern recognition, theorem proving, learning, cognitive science and vision. It also presents introduction to the basic techniques of AI such as heuristic search, semantic nets, production systems, frames, planning and other AI topics. Offered periodically.\n",
	"comp570": "Robotics\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis is a project-oriented course in robotics. Topics are chosen from manipulator motion and control, motion planning, legged-motion, vision, touch sensing, grasping, programming languages for robots and automated factory design. Offered periodically.",
	"comp580": "Database Systems\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics will include relational, hierarchical and network data models; design theory for relational databases and query optimization; classification of data models, data languages; concurrency, integrity, privacy; modeling and measurement of access strategies; and dedicated processors, information retrieval and real time applications. Offered periodically.",
	"comp590": "Computer Architecture\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course is an introduction to the internal structure of digital computers including design of gates, flip-fops, registers and memories to perform operations on numerical and other data represented in binary form; computer system analysis and design; organizational dependence on computations to be performed; and theoretical aspects of parallel and pipeline computation. Offered periodically.",
	"comp594": "Computer Networks\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nThis course provides an introduction to fundamental concepts in computer networks, including their design and implementation. Topics include network architectures and protocols, placing emphasis on protocol used in the Internet; routing; data link layer issues; multimedia networking; network security; and network management. Offered periodically.\n",
	"comp596": "Topics in Computer Science\n(3 credits)\nPrerequisite: Admission to the MS program in Computer Science or consent of instructor\nIn this course, topics are chosen from program verification, formal semantics, formal language theory, concurrent programming, complexity or algorithms, programming language theory, graphics and other computer science topics. Repeatable for different topics. Offered as topics arise.",
	"comp598": " Computer Science Graduate Internship\n(3 credits)\nPrerequisite: Matriculation in the computer science master’s program; at least six credits of graduate-level course work in computer science (COMP); formal application required\nAn internship provides an opportunity to apply what has been learned in the classroom and allows the student to further professional skills. Faculty supervision allows for reflection on the internship experience and connects the applied portion of the academic study to other courses. Repeatable; may earn a maximum of six credits, however, only three credits can be used toward the degree. Graded on (P) Pass/(N) No Pass basis. Offered fall and spring semesters.\n",
}

//go:embed graphics/*
var EmbeddedAssets embed.FS

var counter = 0
var demoApp GuiApp
var textWidget *widget.Text

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowTitle("Minimal Ebiten UI Demo")

	demoApp = GuiApp{AppUI: MakeUIWindow()}

	err := ebiten.RunGame(&demoApp)
	if err != nil {
		log.Fatalln("Error running User Interface Demo", err)
	}
}

func (g GuiApp) Update() error {
	//TODO finish me
	g.AppUI.Update()
	return nil
}

func (g GuiApp) Draw(screen *ebiten.Image) {
	//TODO finish me
	g.AppUI.Draw(screen)
}

func (g GuiApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type GuiApp struct {
	AppUI *ebitenui.UI
}

func MakeUIWindow() (GUIhandler *ebitenui.UI) {
	background := image.NewNineSliceColor(color.Gray16{})
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20))),
		widget.ContainerOpts.BackgroundImage(background))
	textInfo := widget.TextOptions{}.Text("This is our first Window", basicfont.Face7x13, color.White)

	idle, err := loadImageNineSlice("button-idle.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	hover, err := loadImageNineSlice("button-hover.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	pressed, err := loadImageNineSlice("button-pressed.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	disabled, err := loadImageNineSlice("button-disabled.png", 20, 0)
	if err != nil {
		log.Fatalln(err)
	}
	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}
	button := widget.NewButton(
		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),
		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Press Me", basicfont.Face7x13, &widget.ButtonTextColor{
			Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  30,
			Right: 30,
		}),
		// ... click handler, etc. ...
		widget.ButtonOpts.ClickedHandler(FunctionNameHere),
	)
	rootContainer.AddChild(button)
	resources, err := newListResources()
	if err != nil {
		log.Println(err)
	}
	listWidget := widget.List{
		entries: nil,
	}

	textWidget = widget.NewText(textInfo)
	rootContainer.AddChild(textWidget)

	GUIhandler = &ebitenui.UI{Container: rootContainer}
	return GUIhandler
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i := loadPNGImageFromEmbedded(path)

	w, h := i.Size()
	return image.NewNineSlice(i,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

func loadPNGImageFromEmbedded(name string) *ebiten.Image {
	pictNames, err := EmbeddedAssets.ReadDir("graphics")
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := EmbeddedAssets.Open("graphics/" + name)
	if err != nil {
		log.Fatal("failed to load embedded image ", embeddedFile, err)
	}
	rawImage, err := png.Decode(embeddedFile)
	if err != nil {
		log.Fatal("failed to load embedded image ", name, err)
	}
	gameImage := ebiten.NewImageFromImage(rawImage)
	return gameImage
}

func FunctionNameHere(args *widget.ButtonClickedEventArgs) {
	counter++
	message := fmt.Sprintf("You have pressed the button %d times", counter)
	textWidget.Label = message

}
