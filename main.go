package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"portfolio/src"
)

type Media struct {
	Type    string
	URL     string
	OpenURL string
	Alt     string
	Caption string
}

type Project struct {
	Slug            string
	Eyebrow         string
	Title           string
	Meta            string
	Paragraphs      []string
	Tags            []string
	Feature         bool
	RepoURL         string
	SubmissionURL   string
	SubmissionLabel string
	Media           []Media
}

type SiteData struct {
	Name       string
	University string
	Major      string
	Email      string
	GitHubURL  string
	GitHubText string
	SiteURL    string
	Projects   []Project
}

type ProjectPageData struct {
	Site    SiteData
	Project Project
}

func main() {
	data := SiteData{
		Name:       "Maxim Iliev",
		University: "University of Georgia",
		Major:      "Computer Engineering",
		Email:      "maxgoog06@gmail.com",
		GitHubURL:  "https://github.com/MqxS",
		GitHubText: "github.com/MqxS",
		SiteURL:    "https://mqxs.github.io",
		Projects: []Project{
			{
				Slug:    "motor-controller",
				Eyebrow: "Embedded Systems · Power Electronics · Controls",
				Title:   "Custom Three-Phase FOC Motor Controller",
				Meta:    "STM32 · C · KiCad · SPI · USB-C · SVPWM · Oscilloscope",
				Paragraphs: []string{
					"Designed and built a brushless motor controller from scratch after starting with little background in motor theory, PCB design, or field-oriented control. The board uses an STM32 microcontroller and TI gate driver, supports an optional external encoder over SPI, and communicates over USB-C.",
					"Developed the firmware in C, experimented with conventional PWM and SVPWM, explored interpolation techniques for a low-resolution internal Hall encoder, verified PWM timing and deadtime with an oscilloscope, and tuned the controller through iterative hardware testing.",
				},
				Tags:    []string{"FOC", "Motor Control", "Firmware", "PCB Design", "Hardware Debugging", "Control Systems"},
				Feature: true,
				RepoURL: "https://github.com/MqxS/Motor-Controller",
				Media: []Media{
					{Type: "image", URL: "https://raw.githubusercontent.com/MqxS/Motor-Controller/main/Board.jpg", OpenURL: "https://github.com/MqxS/Motor-Controller/blob/main/Board.jpg", Alt: "Custom motor controller PCB", Caption: "Assembled motor controller PCB"},
					{Type: "image", URL: "https://raw.githubusercontent.com/MqxS/Motor-Controller/main/Layout.png", OpenURL: "https://github.com/MqxS/Motor-Controller/blob/main/Layout.png", Alt: "KiCad PCB layout for the custom motor controller", Caption: "KiCad PCB layout"},
					{Type: "video", URL: "../static/assets/motor-controller/Spinning.mp4", OpenURL: "../static/assets/motor-controller/Spinning.mp4", Alt: "Motor controller spinning a brushless motor", Caption: "Controller running a brushless motor"},
				},
			},
			{
				Slug:    "soltix",
				Eyebrow: "Systems Software · Security",
				Title:   "Soltix.cc",
				Meta:    "C · C++ · Go · Windows · Linux",
				Paragraphs: []string{
					"Founded and developed a commercial cheat-detection and diagnostics platform for Minecraft communities. Built low-level systems that analyzed process memory and machine artifacts, then combined those findings into a score for suspected cheating.",
					"Supported many users across various systems, worked through reliability and compatibility issues, maintained backend services, and responded to bad actors attempting to disrupt the platform.",
				},
				Tags: []string{"Process Memory", "Windows Internals", "Backend Systems", "Security"},
			},
			{
				Slug:    "frc-1683",
				Eyebrow: "Robotics · Autonomy",
				Title:   "FRC Team 1683",
				Meta:    "Programming Lead · 2 Years",
				Paragraphs: []string{
					"Spent four years on FRC Team 1683 and served as programming lead for two. Wrote the majority of robot software and supporting vision systems, taught newer students, and developed autonomous and path-following solutions including A*-based approaches.",
					"Also worked extensively across electrical, fabrication, CAD/CAM, and CNC machining, giving me experience integrating software with real mechanical and electrical systems. The team qualified for the FIRST Championship every year I participated.",
				},
				Tags: []string{"Autonomous Systems", "Computer Vision", "Path Planning", "Swerve", "CAD/CAM"},
			},
			{
				Slug:            "sortify",
				Eyebrow:         "AI · Embedded Integration",
				Title:           "Sortify",
				Meta:            "Google Gemini Track Winner · MakeMITxHarvard",
				Paragraphs:      []string{"Designed, built, and programmed a vision-based waste sorting system. The system used AI to classify waste and a compact differential mechanism to redirect items into the correct section of the bin."},
				Tags:            []string{"Gemini API", "Raspberry Pi", "Arduino", "Computer Vision", "C++", "Python"},
				SubmissionURL:   "https://devpost.com/software/sortify-kemb3i",
				SubmissionLabel: "View MakeMITxHarvard submission",
			},
			{
				Slug:            "docdoctor",
				Eyebrow:         "AI · Backend Systems",
				Title:           "DocDoctor",
				Meta:            "MongoDB Track Winner · Google Track Honorable Mention · AI-ATL",
				Paragraphs:      []string{"Built DocDoctor to help customer-support agents find documents more quickly. On the backend, the project used Cobweb, a semantic document-search system, along with Go, Python, JavaScript, Gemini, and MongoDB."},
				Tags:            []string{"Go", "Python", "JavaScript", "Gemini", "MongoDB", "Semantic Search"},
				SubmissionURL:   "https://devpost.com/software/docdoctor",
				SubmissionLabel: "View AI-ATL submission",
			},
			{
				Slug:            "sophi",
				Eyebrow:         "AI · Education",
				Title:           "Sophi",
				Meta:            "NexHacks · Carnegie Mellon",
				Paragraphs:      []string{"Built an AI learning tool that generates instructor-style questions from course material and gives targeted hints to help students practice problem solving with material tailored to their class."},
				Tags:            []string{"AI", "Python", "JavaScript", "Google Gemini", "Education"},
				SubmissionURL:   "https://devpost.com/software/sophia-c3qw4b",
				SubmissionLabel: "View NexHacks submission",
			},
			{
				Slug:            "planetary-analysis",
				Eyebrow:         "Robotics · Computer Vision",
				Title:           "Distributed Planetary Analysis Platform",
				Meta:            "3rd Place · Autonomous Track · RoboTech",
				Paragraphs:      []string{"Developed software for a humanoid robot and autonomous vehicle using inverse kinematics, fiducial detection, pathfinding, PWM motor/servo control, and obstacle avoidance. The autonomous vehicle ran on a Jetson and used fiducial tracking to navigate its environment."},
				Tags:            []string{"Python", "Java", "JavaScript", "Jetson", "Computer Vision", "Kinematics"},
				SubmissionURL:   "https://devpost.com/software/autonomous-distributed-space-exploration",
				SubmissionLabel: "View RoboTech submission",
			},
		},
	}

	if err := os.MkdirAll("docs", 0755); err != nil {
		panic(err)
	}
	if err := copyDir("static", filepath.Join("docs", "static")); err != nil {
		panic(err)
	}

	funcs := template.FuncMap{"add": func(a, b int) int { return a + b }}
	indexTmpl, err := template.New("template.html").Funcs(funcs).ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	projectTmpl, err := template.ParseFiles("project.html")
	if err != nil {
		panic(err)
	}

	if err := render(filepath.Join("docs", "index.html"), indexTmpl, data); err != nil {
		panic(err)
	}
	for _, project := range data.Projects {
		dir := filepath.Join("docs", project.Slug)
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
		if err := render(filepath.Join(dir, "index.html"), projectTmpl, ProjectPageData{Site: data, Project: project}); err != nil {
			panic(err)
		}
	}

	fmt.Printf("Generated homepage + %d project pages in docs/\n", len(data.Projects))

	if err := src.GeneratePDF("docs/index.html", "Portfolio.pdf"); err != nil {
		log.Fatal(err)
	}
}

func render(path string, tmpl *template.Template, data any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, data)
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(target)
		if err != nil {
			return err
		}
		if _, err = io.Copy(out, in); err != nil {
			out.Close()
			return err
		}
		return out.Close()
	})
}
