package digits

//Placeholder type setup commonly used type for digits
type Placeholder [5]string

var zero = Placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = Placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = Placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = Placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = Placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = Placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = Placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}
var seven = Placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = Placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

//Seperator is an array of strings to represent the seperator in the clock
var Seperator = Placeholder{
	"   ",
	" █ ",
	"   ",
	" █ ",
	"   ",
}

//Numbers array contains array elements, each one is a string to represent a number.
var Numbers = [...]Placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}
