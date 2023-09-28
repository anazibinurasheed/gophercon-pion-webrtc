package main

import (
	"github.com/pion/webrtc/v2"
	"github.com/pion/webrtc/v2/pkg/media"
)

func main() {
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})

	if err != nil {
		panic(err)
	}

	offer, err := peerConnection.CreateOffer(nil)

	if err != nil {
		panic(err)
	}

	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	// send Offer to remote PeerConnection via any protocol
	// receive Answer from remote Peerconnection

	answer := webrtc.SessionDescription{}
	err = peerConnection.SetRemoteDescription(answer)

	if err != nil {
		panic(err)
	}

	// Sending Data (Data channels)
	datachannel, err := peerConnection.CreateDataChannel("my-fun-channel", nil)
	if err != nil {
		panic(err)
	}

	datachannel.OnOpen(func() {
		err = datachannel.SendText("hello world !")
		if err != nil {
			panic(err)
		}
	})

	// Send Video

	videoTrack, err := peerConnection.NewTrack(webrtc.DefaultPayloadTypeVP8, 50000, "video", "pion")
	if err != nil {
		panic(err)
	}

	_, err = peerConnection.AddTrack(videoTrack)
	if err != nil {
		panic(err)
	}

	for {
		frame, _, err := ivf.ParseNextFrame()
		if err != nil {
			panic(err)
		}

		err = videoTrack.WriteSample(media.Sample{Data: frame, Samples: 90000})
		if err != nil {
			panic(err)

		}

	}



	peerConnection.OnTrack(func (track *webrtc.Track,receiver *webrtc.RTPReceiver)  {
		if track.Codec().Name == webrtc.Opus {
			for {
				packet,err=track.ReadRTP() 
				if err != nil {
					panic(err)
			}
		}
	})
}
