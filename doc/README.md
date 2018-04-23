# What passes for documentation

There are two commands here that help with generating a password challenge. They do not fully automate the process. There is a step that involves some manual editing.

Note that these were put together as "throw-away scripts" by someone first playing with Go. Development pretty much stopped once they minimally worked.

This documentation was originally intended for Bugcrowd, so that they could generate the challenges. But it is also useful to those participating in the challenges to see exactly how the challenges were generated.

## Step 1: `gentestpwds` to create random passwords with details

`gentestpwds` generates random wordlist type passwords from a list of words given as standard input. It assumes that the wordlist doesn't contain duplicates. It does no sanity checking on its input.

A typical usage might be 

```bash
gentestpwds -n 5 -s 3 -l 3 < AgileWords.txt > secret-2018-03-30.json
```

This will output four three-word passwords as JSON. A single item will look like

```json
{
    "id": "CPU6BAGE",
    "hint": "3 words",
    "prf": "HMAC-SHA256",
    "rounds": 100000,
    "salt": "00bb202b205f064e30f6fae101162a2e",
    "derived": "91976be95cd28e55e580ee9f69a2139202a9b65eabfbbf33c99bc42e3665564d",
    "pwd": "tanbark artistic callus"
}
```
See the file [secret-2018-03-30.json](./secret-2018-03-30.json) for what such output might look like

## Step 2: Manually edit hints and sample status

The next step is to manually edit that output to

1. Modify hints if desired.
2. Indicate which items are samples (and so will have the passwords displayed).

For the files in this directory, that was

```
$ cp secret-2018-03-30.json edited-2018-03-30.json
$ vim edited-2018-03-30.json
```

Typically, you would not make a copy to edit, but would edit the generated output directly. The only reason that a copy is made here is to help document 
each step.

We can edit an item to look like

```json
{
	"id": "CPU6BAGE",
	"hint": "3 words (one of which was new to me)",
	"sample": true,
	"prf": "HMAC-SHA256",
	"rounds": 100000,
	"salt": "00bb202b205f064e30f6fae101162a2e",
	"derived": "91976be95cd28e55e580ee9f69a2139202a9b65eabfbbf33c99bc42e3665564d",
	"pwd": "tanbark artistic callus"
}
```

## Step 3: `chcreator`

In our examples, `chcreator` has very little to do because our source file for it already includes the derived key, salt, and so on. But it could work from sparser source.  

```bash
$ chcreator  < edited-2018-03-30.json > challenge-2018-03-30.json
```

To get similar output, but with the passwords remaining, use the `-p` flag.

```bash
$ chcreator -p  < edited-2018-03-30.json > answers-2018-03-30.json
```

If you just want to check whether the derived keys provided in the source are correct, use the `-t` flag. This can be used to verify whether a submitted answer is correct.

```
$ chcreator -t < answers-2018-03-30.json 
0 bad derived keys out of 5 tested
```


## Next steps

Keep the "answers" file private and publish the "challenge" file. Destroy the "secret" and "edit" files.