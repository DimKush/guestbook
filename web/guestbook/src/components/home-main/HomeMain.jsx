import React, { useEffect } from 'react'
import wlogImg from "../../assets/wlog.png"
import wlistImg from "../../assets/wlist.png"
import weventsImg from "../../assets/wevents.png"

import { FaList } from 'react-icons/fa'
import { FiUsers } from 'react-icons/fi'
import { MdEvent } from 'react-icons/md'
import { AiOutlineAudit } from 'react-icons/ai'
import Clock from 'react-live-clock';
import "./main-style.scss";
import { Redirect } from 'react-router';
import { Link } from 'react-router-dom';

export default function HomeMain({username , setHeaderDescript}) {
	setHeaderDescript("Home");
	const handleListsClick =() => {
		console.log("POOP");
		return <Redirect from="/" to="/lists"/>
	}
	return(
		<div className="home-content-container">
			<div className="username">
				<h1>Hello, {username}. What we will do today at <Clock format={"HH:mm"} ticking={true} timezone={Intl.DateTimeFormat().resolvedOptions().timeZone}/>? </h1>
			</div>
			<Link to="/lists" className="linkButton">
			<div className="bigButton" onClick={handleListsClick} >
				<div className="ButtonLogo"><FaList/></div>
				<div className="ButtonText">All lists</div>
				<div className="bigButtonFooter"></div>
			</div>
			</Link>
			<Link to="/items" className="linkButton">
				<div className="bigButton">
					<div className="ButtonLogo"><MdEvent/> </div>
					<div className="ButtonText">{username} items</div>
					<div className="bigButtonFooter"></div>
				</div>
			</Link>
			<Link to="/audit" className="linkButton">
				<div className="bigButton">
					<div className="ButtonLogo"><AiOutlineAudit/></div>
					<div className="ButtonText">Show audit events</div>
					<div className="bigButtonFooter"></div>
				</div>
			</Link>
			<div className="bigButton">
				<div className="ButtonLogo"><FiUsers/></div>
				<div className="ButtonText">Show Users</div>
				<div className="bigButtonFooter"></div>
			</div>
		
		</div>
);
}