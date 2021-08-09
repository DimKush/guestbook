import React from 'react'
import { Redirect } from 'react-router-dom';
import "./home-style.scss"

export default function Home({isAuth}) {
	console.log("Home = ", isAuth);
	if (!isAuth) {
		console.log(isAuth);
		return <Redirect to="/login"/>
	}

	const handleClickLogout = () => {

	} 

	return (
		<div className="home-base-container">
			<nav>
				<div className="nav-bg"></div>
				<div className="right-side-container">
					<ul>
						<div className="top-system-right-btn">
							<button onClick={handleClickLogout}>Logout</button>
						</div>
					</ul>
				</div>
			</nav>

		</div>
	);
} 