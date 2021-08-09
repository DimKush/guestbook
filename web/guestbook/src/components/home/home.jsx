import React from 'react'
import { Redirect } from 'react-router-dom';


export default function Home({isAuth}) {
	console.log("Home = ", isAuth);
	if (!isAuth) {
		console.log(isAuth);
		return <Redirect to="/login"/>
	}

	return (
		<div className="base-container">
			
		</div>
	);
} 