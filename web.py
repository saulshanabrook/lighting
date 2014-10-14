import sys
import serial
import time

from flask import Flask
from flask import request
from flask import abort, redirect, url_for


PORT = '/dev/tty.usbmodemfd121' # The port my Arduino is on, on my WinXP box.


app = Flask(__name__)

@app.route("/")
def hello():
    return '''
<form action="/dimmer" method="post">
    <input  name="dimmer" />
    <input name="level" />
    <button type="submit">Send your message</button>
</form>

<form action="/text" method="post">
    <input name="text" />
    <button type="submit">Send your message</button>
</form>

'''






def set_dimmer(dimmer, level):
    serial_value = str(dimmer) + "c" + str(level) + "w"
    write_to_arduino(serial_value, PORT)



@app.route('/dimmer', methods=['POST'])
def dimmer_form():
    set_dimmer(request.form['dimmer'], request.form['level'])
    return redirect('/')


def write_to_arduino(value, port):
    # Open a connection to the serial port.  This will reset the Arduino, and
    # make the LED flash once:
    ser = serial.Serial(port)

    # Must given Arduino time to rest.
    # Any time less than this does not seem to work...
    time.sleep(1.5)

    # Now we can start sending data to it:
    written = ser.write(value)
    ser.close()
    print "Bytes Written to port:", written
    print "Value written to port: '%s'"%value


@app.route('/text', methods=['POST'])
def text():
    _, _1, dimmer, _2, level = request.form['text']
    set_dimmer(dimmer, level)

    return redirect('/')

if __name__ == "__main__":
    app.debug = True

    app.run()
