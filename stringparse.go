package main

import(
    "flag"
    "fmt"
    "os"
)

func main() {
    //Subcommands
    countCommand := flag.NewFlagSet("count", flag.ExitOnError)
    listCommand  := flag.NewFlagSet("list", flag.ExitOnError)

    //Count subcommand flag pointers
    countTextPtr := countCommand.String("text", "", "Text to parse (Required)")
    countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|substring}. (Required)")
    countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
    countUniquePtr := countCommand.Bool("unique", false, "Measure unique values of a metric.")

    //List subcommand flag pointers
    listTextPtr := listCommand.String("text", "", "Text to parse. (Required)")
    listMetricPtr := listCommand.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
    listUniquePtr := listCommand.Bool("unique", false, "Measure unique values of a metric.")

    // Verify that a subcommand has been provided
    // os.Arg[0] is the main command
    if(len(os.Args)< 2) {
         fmt.Println("list or count subcommand is required")
         os.Exit(1)
    }

    //Switch on the subcommand
    switch os.Args[1] {
    case "list":
        listCommand.Parse(os.Args[2:])
    case "count":
        countCommand.Parse(os.Args[2:])
    default:
        flag.printDefaults()
        os.Exit(1)
    }



}
