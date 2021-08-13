import React, { useMemo} from 'react'
import './style.scss'
import { useTable, useGlobalFilter, useFilters } from 'react-table'
import MOCK_DATA from './MOCK_DATA.json'
import { COLUMNS } from './columns'
import { AiOutlineSearch } from 'react-icons/ai'

export default function ListsTable({setHeaderDescript}){
	setHeaderDescript("Lists");
	const columns = useMemo(() => COLUMNS , []);
	const data = useMemo(() => MOCK_DATA, []);
	const[sidebar, setSidebar] = React.useState(false);

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
	
	const showSidebar = () => setSidebar(!sidebar);

	const Sidebar = () => {
		return (
			<div className={sidebar ? "SidebarFilter active" : "SidebarFilter" }>
				<button className="searchClick" onClick={showSidebar}><AiOutlineSearch/></button>
				<div className="filters-container">
				{
					headerGroups.map(headerGroup => (
						<div {...headerGroup.getHeaderGroupProps()}>
							{
								headerGroup.headers.map(column => (
									<div {...column.getHeaderProps()} className="search-field">
										{column.render('Header')}
										{ <div>{column.canFilter ? column.render('Filter') : null } </div> }
									</div>
								))
							}
							
						</div>
					))
				}
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