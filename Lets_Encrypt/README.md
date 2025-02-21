I use Let's Encrypt (https://letsencrypt.org/) TLS certificates for my personal sites.
These certs expire after three months, so some time back, I wrote a script to automatically
update DNS for the verification record and obtain a certificate renewal without me doing
anything but running the script.

Every three months, LE would dutifully send me an email reminding me it was soon time to renew my certs—until recently. LE announced it would no longer send out emails to remind folks. Fair enough. So, I set out to write a script to check the expiration and, one week prior, perform the task automatically.

I could have created a scheduled job to run this every 90 days, but I thought this was more fun. I run it once a week.
