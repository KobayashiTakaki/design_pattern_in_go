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
	excecise122()
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

func excecise122() {
	md := NewMultiStringDisplay()

	md.add("Hi!")
	md.add("Good morning.")
	md.add("Good night!")
	md.Show()

	d1 := NewSideBorder(md, "#")
	d1.Show()

	d2 := NewFullBorder(md)
	d2.Show()
}
