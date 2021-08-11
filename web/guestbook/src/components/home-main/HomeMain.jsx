import React, { useEffect } from 'react'
import wlogImg from "../../assets/wlog.png"
import wlistImg from "../../assets/wlist.png"
import weventsImg from "../../assets/wevents.png"
import Clock from 'react-live-clock';

export default function HomeMain({username}) {
	return (
		<div className="home-content-container">
		<div className="username">
			<h1>Hello, {username}. What we will do today at <Clock format={"HH:mm"} ticking={true} timezone={Intl.DateTimeFormat().resolvedOptions().timeZone}/>? </h1>
		</div>
		<div className="bigButton">
			<img className="ButtonLogo" src={wlistImg} alt="Logo"/>
			<div className="ButtonText">All lists</div>
			<div className="bigButtonFooter"></div>
		</div>
		<div className="bigButton">
			<img className="ButtonLogo" src={weventsImg} alt="Logo"/>
			<div className="ButtonText">All events</div>
			<div className="bigButtonFooter"></div>
		</div>
		<div className="bigButton">
			<img className="ButtonLogo" src={wlogImg} alt="Logo"/>
			<div className="ButtonText">Show audit events</div>
			<div className="bigButtonFooter"></div>
		</div>
		
		</div>
);
}