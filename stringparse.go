package main

import(
    "flag"
    "fmt"
    "os"
    "strings"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
    return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
    *s = strings.Split(value, ",")
    return nil
}

func main() {
    //Subcommands
    countCommand := flag.NewFlagSet("count", flag.ExitOnError)
    listCommand  := flag.NewFlagSet("list", flag.ExitOnError)

    //Count subcommand flag pointers
    countTextPtr := countCommand.String("text", "", "Text to parse (Required)")
    countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|substring}. (Required)")
    countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
    countUniquePtr := countCommand.Bool("unique", false, "Measure unique values of a metric.")

    // Use flag.Var to create a flag of our new flagType
    // Default value is the current value at countStringListPtr (currently a nil value)
    var countStringList stringList
    countCommand.Var(&countStringList, "substringList", "A comma seperated list of substrings to be counted.")

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
        flag.PrintDefaults()
        os.Exit(1)
    }

    //Check which subcommand was Parsed using FlagSet
    if listCommand.Parsed() {
        if *listTextPtr == "" {
            listCommand.PrintDefaults()
            os.Exit(1)
        }

        //Choice Flag
        metricChoices := map[string]bool{"chars" : true, "words" : true, "lines" : true}
        if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
            listCommand.PrintDefaults()
            os.Exit(1)
        }
        // Print
        fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n",
            *listTextPtr,
            *listMetricPtr,
            *listUniquePtr,
        )
    }

    if countCommand.Parsed() {
        if *countTextPtr == "" {
            countCommand.PrintDefaults()
            os.Exit(1)
        }

        if *countMetricPtr != "substring" && (*countSubstringPtr != "" || (&countStringList).String() != "[]") {
            fmt.Println("--substring and --substringList may only be used with --metric=substring.")
            countCommand.PrintDefaults()
            os.Exit(1)
        }
        //Choice flag
        metricChoices := map[string]bool{"chars": true, "words": true, "lines": true, "substring": true}
        if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
            countCommand.PrintDefaults()
            os.Exit(1)
        }
        //Print
        fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, substringListPtr: %v, uniquePtr: %t\n",
            *countTextPtr,
            *countMetricPtr,
            *countSubstringPtr,
            (&countStringList).String(),
            *countUniquePtr,
        )
    }
}
