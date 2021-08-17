import React, { Children } from 'react';
import "./modal_style.scss"
import { motion } from "framer-motion"

const ModalLoading = ({active}) => {

	console.log("ACTIVE ? ", active);
	const spinTransition = {
		loop: Infinity,
		duration: 1,
	}

	return (
		<div className={active ? "modal active": "modal "}>
			<div className={active ? "donutContainer active" : "donutContainer"} onClick={e => e.stopPropagation()}>
				<motion.span className="Donut" animate={{rotate: 360 }} transition={spinTransition}/>
			</div>
			
		</div>
	);
};

export default ModalLoading;