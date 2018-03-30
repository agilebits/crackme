# What passes for documentation

There are two commands here that help with generating a password challenge. They do not fully automate the process. There is a step that involves some manual editing.

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

The next step is to manually edit that add to hints, if desired and indicate which items are samples (and so will have the passwords displayed).


