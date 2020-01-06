import React, { useState, useEffect } from 'react';
import { useRouteMatch } from 'react-router-dom';
import NewForm from './NewForm';
import styled from 'styled-components';

let Title = styled.h2`
  color: #036;
  font-weight: normal;
  margin-top: 0;
`;

let Header = styled.header`
  display: flex;
  display: -webkit-flex;
  justify-content: space-between;
  align-items: baseline;
`;

let SettlementTable = styled.table`
　width: 100%;
　border-collapse: collapse;
`;

let Settlement = styled.tr`
  &:nth-child(odd) {
    background: #eef;
  }
`;

let Adjustment = styled(Settlement)`
  color: #009;
`;

let Total = styled.tr`
  background: #ccf;
  font-weight: bold;
`;

let SettlementTitle = styled.td`
  width: 60%;
  padding: 0.5rem 0.7rem;
`;

let SettlementAmount = styled.td`
  width: 15%;
  text-align: right;
  padding: 0.5rem 0.7rem;
`;

let DeleteColumn = styled.td`
  width: 10%;
  text-align: right;
  font-size: smaller;
  padding: 0.5rem 0.7rem;
`;

let DeleteButton = styled.button`
  background: none;
  border: none;
  color: #666;
  &:hover {
    color: #900;
  }
`;

function Monthly() {
    let { params: { year, month } } = useRouteMatch();
    useEffect(() => {
        fetchRecords();
    }, []);

    let [settlements, setSettlements] = useState([]);
    let [adjustment, setAdjustment] = useState({ amounts: [0, 0] });
    let [total, setTotal] = useState({ amounts: [0, 0] });

    function fetchRecords() {
        setSettlements([
            { title: '商品0', amounts: [1000, 0] },
            { title: '商品1', amounts: [0, 2000] },
        ]);
        setAdjustment({
            amounts: [1000, -1000],
        });
        setTotal({
            amounts: [1000, -1000],
        });
    }

    function number(n) {
        if (n === 0) {
            return "";
        } else {
            return n.toString();
        }
    }

    function deleteItem() {
        let ans = window.confirm('Delete the item?');
        if (ans) {

        } else {


        }
    }

    return (
        <div className="Monthly">
            <Header>
                <Title>{year}/{month}</Title>
                <NewForm />
            </Header>
            <SettlementTable>
                <tbody>
                    {settlements.map((s, i) =>
                        <Settlement key={i}>
                            <SettlementTitle>{s.title}</SettlementTitle>
                            {s.amounts.map((a, i) =>
                                <SettlementAmount key={i}>{number(a)}</SettlementAmount>
                            )}
                            <DeleteColumn>
                                <DeleteButton onClick={deleteItem}>Delete</DeleteButton>
                            </DeleteColumn>
                        </Settlement>
                    )}
                    <Adjustment>
                        <SettlementTitle>Adjustment</SettlementTitle>
                        {adjustment.amounts.map((a, i) =>
                            <SettlementAmount key={i}>{a}</SettlementAmount>
                        )}
                        <td></td>
                    </Adjustment>
                    <Total>
                        <SettlementTitle>Total</SettlementTitle>
                        <SettlementAmount>{total.amounts[0]}</SettlementAmount>
                        <SettlementAmount>{total.amounts[1]}</SettlementAmount>
                        <td></td>
                    </Total>
                </tbody>
            </SettlementTable>
        </div>
    );
}

export default Monthly;
