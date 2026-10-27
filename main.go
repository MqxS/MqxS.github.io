package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	Summary         string
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
	Description string
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
	outputDir := flag.String("out", "docs", "Directory for the generated website")
	generatePDF := flag.Bool("pdf", true, "Also generate Portfolio.pdf")
	flag.Parse()

	data := SiteData{
		Name:        "Maxim Iliev",
		Description: "Computer engineering student at the University of Georgia building embedded hardware, robotics, and systems software.",
		University:  "University of Georgia",
		Major:       "Computer Engineering",
		Email:       "maxim.iliev.a@gmail.com",
		GitHubURL:   "https://github.com/MqxS",
		GitHubText:  "github.com/MqxS",
		LinkedInURL: "https://www.linkedin.com/in/maxim-iliev/",
		SiteURL:     "https://mqxs.github.io",
		Projects: []Project{
			{
				Slug:    "motor-controller",
				Eyebrow: "Embedded Systems · Motor Control",
				Title:   "Brushless Motor Controller",
				Summary: "Custom STM32 motor controller with a three phase PCB and field oriented control firmware.",
				Meta:    "C · STM32 · KiCad · SPI",
				Paragraphs: []string{
					"Designed a three phase motor controller PCB in KiCad and wrote C firmware for PWM, SVPWM, and field oriented control. The board has USB-C connectivity and supports an optional high resolution SPI absolute encoder.",
					"Verified PWM timing and deadtime with an oscilloscope, then debugged gate drive and motor control behavior on the assembled board.",
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
				Summary: "Windows process analysis and cheat detection, built with a C/C++ driver and Go tools.",
				Meta:    "Founder & Developer · Jan 2022–Jul 2025",
				Paragraphs: []string{
					"Founded a commercial cheat detection service that analyzed Windows process memory and system artifacts. It reached more than 300 users and 37 customers.",
					"Built a Windows driver in C/C++ and process analysis tools in Go, and administered Linux systems for customer deployments.",
				},
				Tags: []string{"Windows Driver", "Process Memory", "C/C++", "Go", "Linux"},
			},
			{
				Slug:    "frc-1683",
				Eyebrow: "Robotics · Autonomy",
				Title:   "FRC Team 1683",
				Summary: "Autonomous robot software and computer vision for FIRST World Championship competition.",
				Meta:    "Team Member · Aug 2021–May 2025 · Programming Lead for 2 years",
				Paragraphs: []string{
					"Led programming for two of my four years on the team, developing autonomous and computer vision software for competition robots. The team of 85 students reached the FIRST World Championship in both years I served as lead.",
					"Also contributed to robot design, electrical systems, fabrication, and CNC machining.",
				},
				Tags: []string{"Robot Software", "Autonomous Systems", "Computer Vision", "Electrical Systems", "CNC"},
			},
			{
				Slug:            "sidequests",
				Eyebrow:         "Backend Development · HackGT 13",
				Title:           "SideQuests",
				Summary:         "A Go and MongoDB backend for shared outings, activity planning, and group chat.",
				Meta:            "MongoDB Prize · 1st Place, NSA HEARSAY",
				Paragraphs:      []string{"Built the Go backend and server infrastructure for a MongoDB backed app that combines time and location aware activity plans with shared outings, interest matching, and group chat."},
				Tags:            []string{"Go", "MongoDB", "Backend Development", "Server Infrastructure"},
				SubmissionURL:   "https://devpost.com/software/sidequestz",
				SubmissionLabel: "View HackGT submission",
			},
			{
				Slug:            "sortify",
				Eyebrow:         "Computer Vision · Mechanism Design",
				Title:           "Sortify",
				Summary:         "Computer vision and a compact sorting mechanism that route waste to the right bin.",
				Meta:            "Google Gemini Track Winner · MakeMIT x Harvard",
				Paragraphs:      []string{"Developed computer vision software to classify waste and pass the results to sorting controls. Designed a compact differential mechanism to route each item to the correct bin."},
				Tags:            []string{"Computer Vision", "Sorting Controls", "Mechanism Design"},
				SubmissionURL:   "https://devpost.com/software/sortify-kemb3i",
				SubmissionLabel: "View MakeMITxHarvard submission",
			},
			{
				Slug:            "docdoctor",
				Eyebrow:         "Backend Development · Document Search",
				Title:           "DocDoctor",
				Summary:         "Live support call transcripts, agent annotations, and document search powered by Go and MongoDB.",
				Meta:            "MongoDB Track Winner · Google Cloud Honorable Mention · AI ATL",
				Paragraphs:      []string{"Built a Go backend and WebSocket pipeline for live support call transcripts and agent annotations. Added MongoDB persistence and document search APIs so agents could retrieve relevant material during calls."},
				Tags:            []string{"Go", "WebSockets", "MongoDB", "Document Search", "APIs"},
				SubmissionURL:   "https://devpost.com/software/docdoctor",
				SubmissionLabel: "View AI-ATL submission",
			},
			{
				Slug:            "sophi",
				Eyebrow:         "AI · Education",
				Title:           "Sophi",
				Summary:         "An AI study tool that turns course material into practice questions and targeted hints.",
				Meta:            "NexHacks · Carnegie Mellon",
				Paragraphs:      []string{"Built an AI learning tool that generates questions in an instructor's style from course material and gives targeted hints to help students practice problem solving with material tailored to their class."},
				Tags:            []string{"AI", "Learning Tools", "Education"},
				SubmissionURL:   "https://devpost.com/software/sophia-c3qw4b",
				SubmissionLabel: "View NexHacks submission",
			},
			{
				Slug:            "planetary-analysis",
				Eyebrow:         "Robotics · Computer Vision",
				Title:           "Distributed Planetary Analysis Platform",
				Summary:         "Autonomous exploration with Jetson vision, pathfinding, and humanoid robot control.",
				Meta:            "3rd Place · Autonomous Track · RoboTech",
				Paragraphs:      []string{"Developed software for a humanoid robot and autonomous vehicle using inverse kinematics, fiducial detection, pathfinding, PWM motor/servo control, and obstacle avoidance. The autonomous vehicle ran on a Jetson and used fiducial tracking to navigate its environment."},
				Tags:            []string{"Jetson", "Computer Vision", "Inverse Kinematics", "Pathfinding"},
				SubmissionURL:   "https://devpost.com/software/autonomous-distributed-space-exploration",
				SubmissionLabel: "View RoboTech submission",
			},
		},
	}

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		panic(err)
	}
	if err := copyDir("static", filepath.Join(*outputDir, "static")); err != nil {
		panic(err)
	}
	if err := generateSocialPreviews(data, *outputDir); err != nil {
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

	if err := render(filepath.Join(*outputDir, "index.html"), indexTmpl, data); err != nil {
		panic(err)
	}
	for _, project := range data.Projects {
		dir := filepath.Join(*outputDir, project.Slug)
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
		if err := render(filepath.Join(dir, "index.html"), projectTmpl, ProjectPageData{Site: data, Project: project}); err != nil {
			panic(err)
		}
	}

	fmt.Printf("Generated homepage, %d project pages, and %d preview images in %s/\n", len(data.Projects), len(data.Projects)+1, *outputDir)

	if *generatePDF {
		if err := src.GeneratePDF(filepath.Join(*outputDir, "index.html"), "Portfolio.pdf"); err != nil {
			log.Fatal(err)
		}
	}
}

type previewLine struct {
	Text string
	Y    int
}

type socialPreview struct {
	Slug, Name, Heading          string
	TitleLines, DescriptionLines []previewLine
}

func generateSocialPreviews(data SiteData, outputDir string) error {
	socialDir := filepath.Join(outputDir, "static", "assets", "social")
	if err := os.MkdirAll(socialDir, 0755); err != nil {
		return err
	}
	tmpl, err := template.ParseFiles("social.svg")
	if err != nil {
		return err
	}
	previews := []socialPreview{{
		Slug: "portfolio", Name: data.Name, Heading: "ENGINEERING PORTFOLIO",
		TitleLines:       []previewLine{{"I like building things.", 308}},
		DescriptionLines: []previewLine{{"Embedded hardware, robotics,", 379}, {"and systems software.", 417}},
	}}
	for _, project := range data.Projects {
		title := previewLines(project.Title, 23, 308, 66)
		descriptionY := 379
		if len(title) > 1 {
			// Leave room for two headline lines above the description and footer.
			for i := range title {
				title[i].Y -= 26
			}
			descriptionY = 410
		}
		previews = append(previews, socialPreview{
			Slug: project.Slug, Name: data.Name, Heading: strings.ToUpper(project.Eyebrow),
			TitleLines: title, DescriptionLines: previewLines(project.Summary, 58, descriptionY, 38),
		})
	}
	for _, preview := range previews {
		path := filepath.Join(socialDir, preview.Slug)
		if err := render(path+".svg", tmpl, preview); err != nil {
			return err
		}
		if err := src.GenerateSocialImage(path+".svg", path+".png"); err != nil {
			return fmt.Errorf("generate %s preview: %w", preview.Slug, err)
		}
	}
	return nil
}

// previewLines wraps at word boundaries to keep copy inside the preview artwork.
func previewLines(text string, maxChars, startY, lineHeight int) []previewLine {
	var lines []previewLine
	for _, word := range strings.Fields(text) {
		if len(lines) == 0 || len(lines[len(lines)-1].Text)+1+len(word) > maxChars {
			lines = append(lines, previewLine{word, startY + len(lines)*lineHeight})
		} else {
			lines[len(lines)-1].Text += " " + word
		}
	}
	return lines
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
