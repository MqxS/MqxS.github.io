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
	Name        string
	University  string
	Major       string
	Email       string
	GitHubURL   string
	GitHubText  string
	LinkedInURL string
	SiteURL     string
	Projects    []Project
}

type ProjectPageData struct {
	Site    SiteData
	Project Project
}

func main() {
	data := SiteData{
		Name:        "Maxim Iliev",
		University:  "University of Georgia",
		Major:       "Computer Engineering",
		Email:       "maxgoog06@gmail.com",
		GitHubURL:   "https://github.com/MqxS",
		GitHubText:  "github.com/MqxS",
		LinkedInURL: "https://www.linkedin.com/in/maxim-iliev/",
		SiteURL:     "https://mqxs.github.io",
		Projects: []Project{
			{
				Slug:    "motor-controller",
				Eyebrow: "Embedded Systems · Motor Control",
				Title:   "Brushless Motor Controller",
				Meta:    "C · STM32 · KiCad · SPI",
				Paragraphs: []string{
					"Designed a three-phase motor controller PCB in KiCad and wrote C firmware for PWM, SVPWM, and field-oriented control. The board has USB-C connectivity and supports an optional high-resolution SPI absolute encoder.",
					"Verified PWM timing and deadtime with an oscilloscope, then debugged gate-drive and motor-control behavior on the assembled board.",
				},
				Tags:    []string{"Firmware", "PCB Design", "FOC", "SVPWM", "Hardware Debugging"},
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
				Meta:    "Founder · Four-year independent venture",
				Paragraphs: []string{
					"Founded a commercial cheat-detection service that analyzed Windows process memory and system artifacts. It reached more than 300 users and 37 customers.",
					"Built a Windows driver in C/C++ and process-analysis tools in Go, and administered Linux systems for customer deployments.",
				},
				Tags: []string{"Windows Driver", "Process Memory", "C/C++", "Go", "Linux"},
			},
			{
				Slug:    "frc-1683",
				Eyebrow: "Robotics · Autonomy",
				Title:   "FRC Team 1683",
				Meta:    "Programming Lead · 2 Years · Team Member · 4 Years",
				Paragraphs: []string{
					"Led programming for two of my four years on the team, developing autonomous and computer-vision software for competition robots. The 85-student team reached the FIRST World Championship in both years I served as lead.",
					"Also contributed to robot design, electrical systems, fabrication, and CNC machining.",
				},
				Tags: []string{"Robot Software", "Autonomous Systems", "Computer Vision", "Electrical Systems", "CNC"},
			},
			{
				Slug:            "sidequests",
				Eyebrow:         "Backend Development · HackGT 13",
				Title:           "SideQuests",
				Meta:            "MongoDB Prize · 1st Place, NSA HEARSAY",
				Paragraphs:      []string{"Built the Go backend and server infrastructure for a MongoDB-backed app that combines time- and location-aware activity plans with shared outings, interest matching, and group chat."},
				Tags:            []string{"Go", "MongoDB", "Backend Development", "Server Infrastructure"},
				SubmissionURL:   "https://devpost.com/software/sidequestz",
				SubmissionLabel: "View HackGT submission",
			},
			{
				Slug:            "sortify",
				Eyebrow:         "Computer Vision · Mechanism Design",
				Title:           "Sortify",
				Meta:            "Google Gemini Track Winner · MakeMIT x Harvard",
				Paragraphs:      []string{"Developed computer-vision software to classify waste and pass the results to sorting controls. Designed a compact differential mechanism to route each item to the correct bin."},
				Tags:            []string{"Computer Vision", "Sorting Controls", "Mechanism Design"},
				SubmissionURL:   "https://devpost.com/software/sortify-kemb3i",
				SubmissionLabel: "View MakeMITxHarvard submission",
			},
			{
				Slug:            "docdoctor",
				Eyebrow:         "Backend Development · Document Search",
				Title:           "DocDoctor",
				Meta:            "MongoDB Track Winner · Google Cloud Honorable Mention · AI ATL",
				Paragraphs:      []string{"Built a Go backend and WebSocket pipeline for live support-call transcripts and agent annotations. Added MongoDB persistence and document-search APIs so agents could retrieve relevant material during calls."},
				Tags:            []string{"Go", "WebSockets", "MongoDB", "Document Search", "APIs"},
				SubmissionURL:   "https://devpost.com/software/docdoctor",
				SubmissionLabel: "View AI-ATL submission",
			},
			{
				Slug:            "sophi",
				Eyebrow:         "AI · Education",
				Title:           "Sophi",
				Meta:            "NexHacks · Carnegie Mellon",
				Paragraphs:      []string{"Built an AI learning tool that generates instructor-style questions from course material and gives targeted hints to help students practice problem solving with material tailored to their class."},
				Tags:            []string{"AI", "Learning Tools", "Education"},
				SubmissionURL:   "https://devpost.com/software/sophia-c3qw4b",
				SubmissionLabel: "View NexHacks submission",
			},
			{
				Slug:            "planetary-analysis",
				Eyebrow:         "Robotics · Computer Vision",
				Title:           "Distributed Planetary Analysis Platform",
				Meta:            "3rd Place · Autonomous Track · RoboTech",
				Paragraphs:      []string{"Developed software for a humanoid robot and autonomous vehicle using inverse kinematics, fiducial detection, pathfinding, PWM motor/servo control, and obstacle avoidance. The autonomous vehicle ran on a Jetson and used fiducial tracking to navigate its environment."},
				Tags:            []string{"Jetson", "Computer Vision", "Inverse Kinematics", "Pathfinding"},
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
