import React from 'react';
import SignIn from "./sign-in.jsx";
import SignUp from "./sign-up.jsx";

function RightSightComponent({loggingActive, currentState, containerRef, onClick}) {
	return(
	  <div className={loggingActive ? "right-side right" : "right-side left"} ref={containerRef} onClick={onClick}>
		<div className="inner-container">
		   <div className="text">{currentState}</div>
		</div>
	  </div>
	);
}

export default function LoginComponent({isAuth, setAuthStatus}) {
	let current = React.createRef();
	const[isLoggingActive, setLoggingActive] = React.useState(true);
	const[currentState, setCurrentState] = React.useState(!isLoggingActive ? "Sign in" : "Sign up");
	
	const[Username, setUsername] = React.useState("");

	const changeState = () => {
		setLoggingActive(!isLoggingActive);
		setCurrentState(isLoggingActive ? "Sign in" : "Sign up");
	}

	return(
		<div className="form-signin">
			<div className="login">
				<div className="container">
				{isLoggingActive && <SignIn containerRef={(ref) => current = ref} isAuth={isAuth} setAuthStatus={setAuthStatus} />}
				{!isLoggingActive && <SignUp containerRef={(ref) => current = ref} setLoggingActive={setLoggingActive} />}
				</div>
				<RightSightComponent loggingActive={isLoggingActive} currentState={currentState} containerRef={ref => current = ref} onClick={changeState}/>
			</div>
		</div>
	);
}
