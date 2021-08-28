import React, { useState } from 'react'
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";


export default function DeleteList(selectedRow, currentUser) {
	let result = {Status : "", Message : ""};

	if(currentUser !== selectedRow.owner){
		result.Status = "Error";
		result.Message = `Access denied. You don't have permissions to delete with list. Owner of the record is ${selectedRow.owner}.`;

		return result;
	}
	
	(async () => {
		
	})();


	return "OK";
}