# Count words

These are examples of a programs that read the files in a directory and count the number of words in these files.

## Communication over sharing

In the [communication-over-sharing](./communication-over-sharing/) folder there is an example of using a "communication over sharing" approach.

After building the executable, as specified [here](./communication-over-sharing/readme.md) you can read the number of words with the following command

`./bin/countWordsWithCommunication -dir testdata/divina-commedia`

The command `./bin/countWordsWithCommunication` has other parameters that can be seen using the `-h` options like this

`./bin/countWordsWithCommunication -h`

## Sharing over Communication

In the [sharing-over-communication](./sharing-over-communication/) folder there is an example of using a "sharing over communication" approach.

After building the executable, as specified [here](./sharing-over-communication/readme.md) you can read the number of words with the following command

`./bin/countWordsWithSharing -dir testdata/divina-commedia`

The command `./bin/countWordsWithCommunication` has other parameters that can be seen using the `-h` options like this

`./bin/countWordsWithSharing -h`
