import logo from './logo.svg';
import React from 'react';
import './App.scss';
//import { SignIn, SignUp } from "./components/login/index"
import SignIn from "./components/login/sign-in.jsx"
import SignUp from "./components/login/sign-up.jsx"

function RightSightComponent({loggingActive, currentState, containerRef, onClick}) {
  //setCurState(!currentState);
  
  return (
    <div className={loggingActive ? "right-side right" : "right-side left"} ref={containerRef} onClick={onClick}>
      <div className="inner-container">
         <div className="text">{currentState}</div>
      </div>
    </div>
  );
}

export default function App() {
  let current = React.createRef();
  const[isLoggingActive, setLoggingActive] = React.useState(true);
  const[currentState, setCurrentState] = React.useState(!isLoggingActive ? "Sign in" : "Sign up");
  
  const changeState = () => {
    setLoggingActive(!isLoggingActive);
    setCurrentState(isLoggingActive ? "Sign in" : "Sign up");
  }

  return (
    <div className="App">
      <div className="login">
        <div className="container">
          {isLoggingActive && <SignIn containerRef={(ref) => current = ref} />}
          {!isLoggingActive && <SignUp containerRef={(ref) => current = ref} />}
        </div>
        <RightSightComponent loggingActive={isLoggingActive} currentState={currentState} containerRef={ref => current = ref} onClick={changeState}/>
      </div>
    </div>
  );
}

/*
const RightSightComponent = props => {
  return (
    <div className="right-side" ref={props.containerRef} onClick={props.onClick}>
      <div className="inner-container">
         <div className="text">{props.current}</div>
      </div>
    </div>
  );
}

class App extends React.Component{
  constructor(props){
    super(props);
    this.state = {
      isLoggingActive: true
    }
  }  

  changeState() {
    const { isLoggingActive } = this.state;
    if(isLoggingActive) {
      this.LeftSightComponent.classList.remove("right");
      this.LeftSightComponent.classList.add("left");
    } else {
      this.LeftSightComponent.classList.remove("left");
      this.LeftSightComponent.classList.add("right");
    }

    this.setState((prevState) => ({ isLoggingActive : !prevState.isLoggingActive }))
  }

  render(){
    const {isLoggingActive} = this.state;
    let currentState = isLoggingActive ? "Sign up" : "Sign in";
    //const  currentActive = isLoggingActive ? "Sign in" : "Sign up";
    return (
      <div className="App">
        <div className="login">
          <div className="container">
            {isLoggingActive && <SignIn containerRef={(ref) => this.current = ref} />}
            {!isLoggingActive && <SignUp containerRef={(ref) => this.current = ref} />}
          </div>
          <RightSightComponent current={currentState} containerRef={ref => this.LeftSightComponent = ref} onClick={this.changeState.bind(this)}/>
        </div>
      </div>
    );
  }
}

*/