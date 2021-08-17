import React, { Children } from 'react';
import "./modal_style.scss"


const ModalLoading = ({active, setActive}) => {
	return (
		<div className={active ? "modal active": "modal "} onClick={() => setActive(false)}>
			<div className={active ? "modal_thx_content active" : "modal_thx_content"} onClick={e => e.stopPropagation()}>
			</div>
			
		</div>
	);
};

export default ModalLoading;