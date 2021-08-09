import React from 'react'
import { Redirect } from 'react-router-dom';
import "./home-style.scss"

export default function Home({isAuth}) {
	console.log("Home = ", isAuth);
	if (!isAuth) {
		console.log(isAuth);
		return <Redirect to="/login"/>
	}

	return (
		<div className="home-base-container">
			<nav>
				<div class="nav-bg"></div>
				<ul>
					<li><a href="">Link</a></li>
				</ul>
			</nav>

		</div>
	);
} 