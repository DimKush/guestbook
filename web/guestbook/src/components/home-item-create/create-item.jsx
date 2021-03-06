import React, { useEffect, useState } from 'react'
import "./create-item-style.scss"
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";
import { Redirect, Link, useParams} from 'react-router-dom';
import { AiOutlineDoubleLeft } from 'react-icons/ai'

function ColumnCreateList ( {column, ref_current, blocked=false, value} ) {
	return(
		<div className="form-group">
		<span>{column}</span>
			
			<input class="form-field"
				type="text"
				id={column}
				ref={ref_current}
				disabled={blocked}
				defaultValue={value}
				/>
		</div>
	);
}


export default function CreateItem() {
	const[idCheckboxBlocked, setIdCheckboxBlocked]= useState(true);
	const[ownerCheckboxBlocked, setOwnerCheckboxBlocked] = useState(false);
	const[ItemsTypes, setItemsTypes] = useState([]);
	const[currentUser, setCurrentUser] = useState("");
	
	const[modalMsgHead, setModalMsgHead] = React.useState("");
	const[modalMsg, setModalMsg] = React.useState("");
	const[modalActive, setModalActive] = React.useState(false);
	const[listOwner, setListOwner] = useState("");

	let idItemInput = React.createRef();
	let itemTypeInput = React.createRef();
	let descriptionInput = React.createRef();
	let auto_id_checkbox = React.createRef();

	let { id } = useParams();

	useEffect(() => {
		(
		  async () => {
			const responce = await fetch("http://localhost:8007/api/lists/items/types", {
			  headers : {"Content-type" : "application/json",
						 "Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			  
			});
			const content = await responce.json();
			if(content.Status === "OK"){
				setItemsTypes(content.Result);
			} else {
				// TODO : ERROR!
			}
		  }
		)();
		
		(async () => {
			const responce = await fetch("http://localhost:8007/auth/user", {
			  headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
			  credentials : "include",
			});

			const content = await responce.json();

			if(content.Status === "OK"){
				setCurrentUser(content.username);
			} else {
				// TODO: Modal error
			}
		})();
	}, []);

	const handleCreateClick = () => {
		const CreateObjItems = {
			"id" : Number(idItemInput.current.value),
			"description" : descriptionInput.current.value,
			"item_type_name" : itemTypeInput.current.value,//,
		};
		

		(
			async() => {
				const responce = await fetch(`http://localhost:8007/api/lists/${id}/items/create`, {
				method : "POST",
				headers : { "Content-type" : "application/json",
							"Authorization" :`Bearer ${cookies.get("jwt")}`},
				credentials : "include",
				body : JSON.stringify(CreateObjItems),
				});
				const content = await responce.json();
				if(content.Status === "OK"){
					setModalMsg("Record was created.");
					setModalMsgHead("OK");
					setModalActive(true);
				} else {
					setModalMsg(content.Message);
					setModalMsgHead(content.Status);
					setModalActive(true);
				}
			}
		)();
	}

	const handleCleanFieldsClick = () => {
		idItemInput.current.value = "";
		descriptionInput.current.value = "";
		setIdCheckboxBlocked(true);
		auto_id_checkbox.current.checked = true; 

		itemTypeInput.current.selectedIndex = 0;

	}

	return(
		<div className="list-card-main">
			<div className="row-form">
				<h1>Create new Item</h1>
			</div>
			<div className="row-form">
				<div className="search-field id checkbox">
					<input type="checkbox" id="autoId" defaultChecked={idCheckboxBlocked} ref={auto_id_checkbox} onChange={() =>
						{
							setIdCheckboxBlocked(!idCheckboxBlocked)
							idItemInput.current.value = "";
						}}>
						</input>
					<label for ="autoId">Auto-increment item Id</label>
				</div>
				<div className="search-field id">
					<ColumnCreateList column={"Item Id"} ref_current={idItemInput} blocked={idCheckboxBlocked}/>
				</div>
				<div className="search-field id">
					<ColumnCreateList column={"List Id"} blocked={true} value={id}/>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field owner" >
				<div className="form-group">
					<span>Item Type</span>
						<select className="form-field ownerSelect" ref={itemTypeInput} disabled={ownerCheckboxBlocked} onChange={e => setListOwner(e.target.value)}> 
							<option disabled selected value>-- Select an item type --</option>
								{
									ItemsTypes.map(element => (
										<option value={element.fullname}>
											{element.fullname}
										</option>
									))
								}
						</select>
					</div>
				</div>
			</div>
			<div className="row-form">
				<div className="search-field description">
					<ColumnCreateList column={"Description"} ref_current={descriptionInput}/>
				</div>
			</div>
			<div className="row-form">
				<Link to={`/lists/${id}/items`}>
						<button className="control-but back"><AiOutlineDoubleLeft/><div className="but-tab-hight-text">Back</div></button>
				</Link>
				<button className="control-but" onClick={handleCreateClick}>Create Item</button>
				<button className="control-but" onClick={handleCleanFieldsClick}>Clean fields</button>
			</div>
			<Modal active={modalActive} setActive={setModalActive} head={modalMsgHead} msg={modalMsg} isError={false}/>
		</div>
	);
} 