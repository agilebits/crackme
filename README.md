# Password cracking challenges

Our (1Password's) goals in offering these challenges is to gain a better sense of the resistance of various types of user Master Passwords to cracking if 1Password data is captured from a user's device.

## Background

Our use of Two Secret Key Derivation (2SKD) protects users from Master Password cracking attempts in the event that data is captured from our servers, but 2SKD does not offer that protection if data is captured from the user's own device. Thus the strength of user Master Passwords remains an important part of user security for 1Password.

We need to encourage users to use Master Passwords which

1. people can remember
2. people can reasonably enter on their devices
3. are sufficiently strong

We are creating these challenges to help us better understand (3).

## Challenges can't be too hard (or too easy)

We would love for people to use 1Password Master Passwords that are simply too hard to crack in the event that data is captured from their local devices. But if we present cracking challenges that are too hard to win, nobody will take the challenge. Instead, we are offering what we hope are winnable challenges with sufficient prizes that many of them will be won.

Let us emphasize this point for when results come in: **These challenges are intended to be winnable. A success does not indicate any weakness in 1Password.**

This means that the passwords we present here are weaker than we recommend as 1Password Master Passwords. The prizes we offer should be worth the effort that the participants need to put in.

At the same time we want the attempts to take some real effort so that we can get more data on that effort. In any cracking effort, there are some fixed costs of simply setting up the cracking run (preparing the data, configuring the software, etc), we want those costs to be dominated by the actual cracking.

## How we help the participants

Our interest is to understand cracking efforts in terms of the strength of a test Master Password under the assumption that an attacker fully knows the details of key derivation and password generation scheme. Therefore we try to provide everything a participant will need to know to set up their systems prior to the beginning of the competition. Thus we make available

1. The [source for the scripts](https://github.com/agilebits/crackme) used to generate the challenge passwords
2. [Sample challenges](https://github.com/agilebits/crackme/tree/master/doc) (some with "answers") published prior to the official challenge.
3. The KDF we use for these challenges is stripped of many of the idiosyncracies of the 1Password KDF that are not relevant for the difficulty of cracking locally captured data.

Individual challenges will look something like this, but see the source for generating them and the sample docs for more detail.

```
{
    "id": "aXw39qx7a5kt",
    "hint": "3 words",
    "prf": "HMAC-SHA256",
    "rounds": 100000,
    "salt": "697c37f6ac7b6b992d12c8eab3269af6",
    "derived": "3e0f1903cc73b07a7070a661f8450d495cc99151ae67bcdf69a80d0391e7d62f"
}
```

## Administration

To ensure fair handling of the contest itself and the award of payments, we are asking Bugcrowd to administer this.

The Bugcrowd program is at https://bugcrowd.com/onepasswordgame

This is a natural choice, as they both have the experience with delivering bounties, and have earned a reputation as a trusted party in dealing both with those offering bounties and those seeking them.


## Prizes

On June 11, 2018 we doubled our initial prizes and added a fourth place prize. On July 26, we increased the prizes yet again.

- For the first person or team to crack a three word password, we offer ~~4096~~ ~~8192~~ 12288 USD.

- For the second person or team to crack a _different_ three word password, we offer ~~2048~~ ~~4096~~ 8192 USD.

- For the third person or team to crack _yet a different_ three word password, we offer ~~1024~~ ~~2048~~ 6144 USD.

- And for the fourth person or team to crack yet another one, we we offer ~~1024~~ 4096 USD.

If no correct submission has been submitted within one month, we _may_ increasing the prizes. [Update: On June 11 we have increased the prizes. On July 26, we increased the prizes yet again.]

## Additional hints

On August 23, 2018, we published an updated challenge file with 1 bit hints. This lists the first bit of the unsalted SHA256 hash of the solution.

On September 24, 2018, we published an updated challenge file with 2 bit hits. This lists the first two bits of the unsalted SHA256 hash of the solution.

The file with these hints is at 
The hints will be initial bit(s) of a SHA256 unsalted hash of the password.
The nature of the hints has been under discussion in [this forum topic](https://discussions.agilebits.com/discussion/89318/world-password-day-cracking-challenge/p1).

Description | File | Signature file |
-----------------|:----------------------------------------|:-----------------------------------
1 bit hints | [password-day-2018-1bitHints.json](./password-day-2018-1bitHints.json) | [password-day-2018-1bitHints.json.asc](./password-day-2018-1bitHints.json.asc) |
2 bit hints | [password-day-2018-2bit-hints.json](./password-day-2018-2bit-hints.json) | [password-day-2018-2bit-hints.json.asc](./password-day-2018-2bit-hints.json.asc) |


## Rules

1. No one at AgileBits or BugCrowd can win.
2. Social engineering, or gaining the solutions through penetration is not allowed. This is a cracking-only exercise.
3. Participants may only use systems with the owner's permission. You may not steal computing resources in your cracking efforts.
2. Winners must provide a write-up of what they did, with estimations of total cost to crack, guesses per second, the systems used. The write-up need not be submitted at the same time as a successful crack, which need only include the ID of the particular challenge and the successful password.
3. Submission is exclusively [via Bugcrowd](https://bugcrowd.com/onepasswordgame), and Bugcrowd standard rules apply

## Schedule

Announcement of forthcoming challenge (when this and associated documents are published): TBA

Start of contest. Publication of the actual challenges: Noon, EDT on World Password Day, May 3, 2018. 2018-05-03 16:00:00 +0000 UTC

The challenge was published at that time at https://github.com/agilebits/crackme/tree/master/password-day-2018.json

Note that a file with 1 bit hints is at 
https://github.com/agilebits/crackme/tree/master/password-day-2018-1bitHints.json

## Contest status

The game is afoot.

Last update: <!-- date -u "+%Y-%m-%d:%H:%M:%S UTC" --> 2018-11-15:22:34:22 UTC

ID        |  Status   |  Successful password      | Hint | Submission date | By whom       | Place | Write-up location
:---------|-----------|:--------------------------|:-----|-----------------|:--------------|-------|:-----------------
3UOKUEBO  | Sample    | governor washout beak     | 0b01 | N/A             | Sample        | 0th   | N/A
AJPYJUTN  | Sample    | glassy ubiquity absence   | 0b11 | N/A             | Sample        | 0th   | N/A
IV2DL67Q  | Sample    | splendor excel rarefy     | 0b01 | N/A             | Sample        | 0th   | N/A
NO4VRU4S  | Not found |                           | 0b10 |                 |               | Nth   |
33YRS77A  | Not found |                           | 0b01 |                 |               | Nth   |
J6J4QUWQ  | Not found |                           | 0b01 |                 |               | Nth   |
SFELTO3W  | Found     | faint bust perturb        | 0b00 | 2018-11-07      | [List](./write-ups/SFELTO3W.md#members)      | 2nd   | [write up](./write-ups/SFELTO3W.md)
DOHB6DC7  | Found     | mansard humpback unbutton | 0b00 | 2018-10-14      | [List](./write-ups/DOHB6DC7.md#members)      | 1st   | [write up](./write-ups/DOHB6DC7.md)
2SB5OP3G  | Found     | befell car granary        | 0b00 | 2018-11-10      | TBA           | 3rd   | 
5BSLBTKR  | Found     | minute judd obedient      | 0b10 | 2019-01-10      | TBA           | 4th  |
