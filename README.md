slotted-goloha
==============

Implementation of the Slotted Aloha Protocol in Go

If you have data to send, send the data
If the message collides with another transmission, try resending "later"
on collision, sender waits random time before trying again


All frames have the same length.
Stations cannot generate a frame while transmitting or trying to transmit. (That is, if a station keeps trying to send a frame, it cannot be allowed to generate more frames to send.)
The population of stations attempts to transmit (both new frames and old frames that collided) according to a Poisson distribution.
