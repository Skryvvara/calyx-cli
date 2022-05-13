<div align=center>
  <h1>Calyx-Cli</h1>
  <p>a humble cli for me, darkyne.</p>
</div>

## description

Calyx is a cli written in [go][golang] mostly to automate things I have to do and learn [go][golang] along the way. While this is open-source it is tailored to my needs and might not work for everyone so please keep that in mind.

## commands

### create [options]

options:

- --file value, -f value | name of the script file
- --path value, -f path | path of the output file
- --help, -h | show help (default: false)
  > Command to generate files defined by yaml files.
  > example usage: `calyx create --file makehelm --path ./makefile` <br>
  > ouput: `Creates new file called makefile`

### generate [options]

options: <br>

- --bytes value, -b value | set the amount of bytes (default: 32) <br>
- --copy, -c | copy the output string to clipbard (default: false) <br>
- --help, -h | show help (default: false) <br>
  > this command generates a key by giving it an amount of bytes as options. Defaults to 32 bytes.
  > example usage: `calyx generate --copy --bytes 32` <br>
  > output: `342b43a03a9f8aedb2a1f568e019249a6a1cf654e5dbaba9bf4410435a33eac5`

<!-- VARIABLES -->

[golang]: https://go.dev/
