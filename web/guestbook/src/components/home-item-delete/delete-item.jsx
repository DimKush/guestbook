import React, { useState } from 'react'
import { cookies } from "../../App";
import Modal from "../modal/modal.jsx";

async function DropItemBack(list_id, item_id){ 
	console.log("list_id", list_id);
	console.log("item_id", item_id);
	const responce = await fetch(`http://localhost:8007/api/lists/${list_id}/items/${item_id}`, {
		method : "DELETE",
		headers : { "Content-type" : "application/json",
						"Authorization" :`Bearer ${cookies.get("jwt")}`},
		credentials : "include",
	});

	const content = await responce.json();

	return {Status : content.Status, Message : content.Message}
}

export default async function DeleteItem(list_id, selectedRow) {
	let result = {Status : "", Message : ""};

	console.log(list_id , selectedRow);

	result = await DropItemBack(list_id, selectedRow.id);

	if (result.Status !== "OK") {
		result.Message=`Cannot delete item. Backend problem : ${result.Message}`;
	} else {
		result.Message=`Item with id ${selectedRow.id} has been deleted.`
	}

	return result;
}