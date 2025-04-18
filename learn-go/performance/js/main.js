function loop() {
    var count = 0;
    var startTime = Date.now();
    
    while (Date.now() - startTime < (1000 * 10)) { // loop for 1 second
        count++;
        // console.log("JS: looped: ", count)
    }

    console.log("Done", count);
}

loop()