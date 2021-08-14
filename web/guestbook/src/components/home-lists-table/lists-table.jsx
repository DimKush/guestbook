import React, { useMemo} from 'react'
import './style.scss'
import { useTable, useGlobalFilter, useFilters } from 'react-table'
import MOCK_DATA from './MOCK_DATA.json'
import { COLUMNS } from './columns'
import { AiOutlineSearch } from 'react-icons/ai'
import { BsBoxArrowInRight } from "react-icons/bs";

import "./filters-styles.scss"


export default function ListsTable({setHeaderDescript}){
	setHeaderDescript("Lists");
	const columns = useMemo(() => COLUMNS , []);
	const data = useMemo(() => MOCK_DATA, []);
	const[sidebar, setSidebar] = React.useState(false);
	const[clearInput, setClearInput] = React.useState(false);
	let columnsFiltered = {};

	const {
		getTableProps,
		getTableBodyProps,
		headerGroups,
		rows,
		prepareRow,
		state,
	} =  useTable(
		{
			columns,
			data,
		},
		useFilters,
	);
	
	

	function ColumnFilter ( {column} ) {
		console.log("columnsFiltered[column.id]", columnsFiltered[column.id])
		return (
			<div className="form-group">
			<span>{column.id}</span>
				
				<input class="form-field"	
					onChange ={(event) =>  {
						columnsFiltered[column.id] = event.target.value;
					}} />
			</div>
		);
	}
	
	const showSidebar = () => setSidebar(!sidebar);
	const handleClickRefresh = () => {
		console.log("handleClickRefresh");
	}

	const handleClickFind = () => {
		console.log(columnsFiltered);
		headerGroups.map(headerGroup => { headerGroup.headers.map(column =>{
			column.setFilter(columnsFiltered[column.id]);
		} ) })

		console.log("columnsFiltered after ", columnsFiltered);
	}

	const Sidebar = () => {
		return (
			<div className={ sidebar ? "SidebarFilter active" : "SidebarFilter" }>
				{!sidebar && <button className="searchClick" onClick={showSidebar}><AiOutlineSearch/></button>}
				<div className="filters-container">
				{sidebar && 
					<div className="searchButton" onClick={showSidebar}>
					<button className="searchClick active "><BsBoxArrowInRight/> </button>
						<div className="searchText">Search in table </div>
					</div>
				}
				{
					
					headerGroups.map(headerGroup => (
						<div {...headerGroup.getHeaderGroupProps()}>
							{
								headerGroup.headers.map(column => (
									<div {...column.getHeaderProps()} className="search-field">
										{column.canFilter ? <ColumnFilter column={column}/> : null }
									</div>
								))
							}
							
						</div>
					))
				}
				<div className="buttons-place">
					<button className="sidebar-but" onClick={handleClickRefresh}>Refresh</button>
					<button className="sidebar-but" onClick={handleClickFind}>Find</button>
				</div>

				</div>
			</div>

		);
	} 

	return(
	<div className="form-container">
		<Sidebar/>
	<div className={sidebar ? "form-events active" : "form-events"}>
		
		{/* <div className="search-container">
			
		</div> */}
		<table {...getTableProps()} > 
			<thead>
				{
					headerGroups.map(headerGroup => (
						<tr {...headerGroup.getHeaderGroupProps()}>
							{
								headerGroup.headers.map(column => (
									<th {...column.getHeaderProps()}>{column.render('Header')}
										{/* <div>{column.canFilter ? column.render('Filter') : null } </div> */}
									</th>
								))
							}
							
						</tr>
					))
				}

			</thead>
			<tbody {... getTableBodyProps()}>
				{
					rows.map(row => {
						prepareRow(row)
						return (
							<tr {...row.getRowProps()}>
								{
									row.cells.map((cell) => {
										return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
									})
								}
							</tr>
						)
					})
				}
				
			</tbody>
		</table>
	</div>
	</div>
	);
}