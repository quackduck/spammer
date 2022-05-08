# spammer
Need an overpowered, automated spammer CLI to impress / unimpress your friends? This is it!

## Install
```shell
git clone https://github.com/quackduck/spammer
cd spammer
go install
```

## Usage

```text
spammer <message> <repetitions> <delay>
```
```shell
spammer hello 100 10 # spam 'hello' 100 times with a 10ms delay between each run
spammer "Happy Mother's Day!" 1000 0 # wish your mom 1000 times as fast as possible
```

When you run one of these, you'll see:

```text
starting in 10 seconds, go to the app to spam, hit q three times in a second to quit
```

Just like it says, you should go to the app you want to spam and wait 10 seconds for it to start (the buffer is so you have time to switch).

If you feel like you should stop spamming / you're friends start threatening you, just hit <kbd>Q</kbd> thrice to stop.

I've been able to send 100 messages on my M1 mac air in ~1.4 seconds when running at full speed 😎:
![image](https://user-images.githubusercontent.com/38882631/167321090-6913033e-0b30-4f86-b499-931cdc1caf8f.png)


## Fancy legal disclaimer
_Ishan Goel is not responsible for any bad things that happen to you or other people as a result of you using this._
