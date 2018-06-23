package fuzz

const longHelpText = `
Use the file filenames.txt as input, hide all 200 and 404 responses:

    monsoon fuzz --file filenames.txt \
      --hide-status 200,404 \
      https://example.com/FUZZ

Skip the first 23 entries in filenames.txt and send at most 2000 requests:

    monsoon fuzz --file filenames.txt \
      --skip 23 \
      --limit 2000 \
      --hide-status 404 \
      https://example.com/FUZZ

Hide responses with body size between 100 and 200 bytes (inclusive), exactly
533 bytes or more than 10000 bytes:

    monsoon fuzz --file filenames.txt \
      --hide-body-size 100-200,533,10000- \
      https://example.com/FUZZ

Try all strings in passwords.txt as the password for the admin user, using an
HTTP POST request:

    monsoon fuzz --file passwords.txt \
      --method POST \
      --data 'username=admin&password=FUZZ' \
      --hide-status 403 \
      https://example.com/login

Run requests with a range from 100 to 500 with the request value inserted into
the cookie "sessionid":

    monsoon fuzz --range 100-500 \
      --header 'Cookie: sessionid=FUZZ' \
      --hide-status 500 https://example.com/login/session

Request 500 session IDs and extract the cookie values (matching case insensitive):

    monsoon fuzz --range 1-500 \
      --extract '(?i)Set-Cookie: (.*)' \
      https://example.com/login

Hide responses which contain a Date header with an uneven number of seconds:

    monsoon fuzz --range 1-500 \
      --hide-pattern 'Date: .* ..:..:.[13579] GMT' \
      https://example.com/FUZZ

Only show responses which contain the pattern "The secret is: " in the response:

    monsoon fuzz --range 1-500 \
      --show-pattern 'The secret is: ' \
      https://example.com/FUZZ


Filter Evaluation Order
#######################

The filters are evaluated in the following order. A response is displayed if:

 * The status code is not hidden (--hide-status)
 * The header and body size are not hidden (--header-size, --body-size)
 * The header and body does not contain a hide pattern (--hide-pattern)
 * The header or body contain all show pattern (--show-pattern, if specified)


References
##########

The regular expression syntax documentation can be found here:
https://golang.org/pkg/regexp/syntax/#hdr-Syntax
`
