package main

func main() {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1, "#")
	b3 := NewFullBorder(b2)
	b1.Show()
	b2.Show()
	b3.Show()
	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewStringDisplay("Hello, world"),
					"*",
				),
			),
		),
		"/",
	)
	b4.Show()
	excecise121()
}

func excecise121() {
	b1 := NewStringDisplay("Hello, world.")
	b2 := NewUpDownBorder(b1, "-")
	b3 := NewSideBorder(b2, "*")
	b1.Show()
	b2.Show()
	b3.Show()
	b4 := NewFullBorder(
		NewUpDownBorder(
			NewSideBorder(
				NewUpDownBorder(
					NewSideBorder(
						NewStringDisplay("Hello, world."),
						"*",
					),
					"=",
				),
				"|",
			),
			"/",
		),
	)
	b4.Show()
}
