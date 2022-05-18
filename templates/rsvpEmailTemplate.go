package templates

const RsvpHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>RsvpUser</title>
</head>
<body>
    <h1>Thank you for RSVP to this event</h1>
    <p>Your Personal QR CODE: </p>
    <img width="256" height="256" src="{{.QrLink}}" alt="QR-CODE">
    <p>Please donot share this QR with anyone</p>
</body>
</html>
`