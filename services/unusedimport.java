import java.util.ArrayList;    // Used
import java.util.List;         // Used
import java.util.HashSet;      // Unused
import java.util.Date;         // Unused
import java.io.*;              // Wildcard (not flagged)
import static java.util.Collections.emptyList; // Static import (unused)

public class SampleClass {
    public static void main(String[] args) {
        // Use List/ArrayList
        List<String> items = new ArrayList<>();
        items.add("Test");
        System.out.println(items);

        // Date/HashSet never used
        // emptyList() never used
    }
}