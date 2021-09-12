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

## Verifying submission

Submissions are through the Bugcrowd submission process for the https://bugcrowd.com/onepasswordgame program. Timestamps will be from the Bugcrowd tracker.

### Is the submission valid?

1. Is the submission via the Bugcrowd program?
1. Does the submission contain a challenge ID and a single candidate password?

### Is the candidate password correct?

2. Verify that the challenge ID is among the challenges.
3. Copy the corresponding entry into a file, say `candidate-ID-Submitter.json`
4. Add the candidate password as a `"pwd"` field to that challenge.
5. run `chcreator -t < candidate-ID-Submitter.json`
6. If all goes well, it should return "0 bad derived keys out of 1 tested"
7. If it reports a bad derived key, then double check for any editing/coping errors. Perhaps test your procedures on the item listed above in this document. Unless there is reason to believe that the submission is insincere, ask someone else to check as well.

### Handling a correct submission

We need to let all participants know that the particular item has been cracked

1. Modify to brief to list the that item ID as found, and what the password is.
2. Let the submitters know that their submission was correct. Remind them that to collect the prize they will need to submit a write up, describing the resources used, guessing rates, etc.
3. Tweet, or otherwise, publicize the find beyond the brief.

### Handling incorrect submissions

This really shouldn't happen. Once a candidate is considered, it should be easy to determine whether it works or not. We will deal with this situation if it arises.
