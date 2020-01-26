import React, { useState, useEffect } from 'react';
import { useRouteMatch } from 'react-router-dom';
import NewForm from './NewForm';
import styled from 'styled-components';

const Title = styled.h2`
  color: #036;
  font-weight: normal;
  margin-top: 0;
`;

const Header = styled.header`
  display: flex;
  display: -webkit-flex;
  justify-content: space-between;
  align-items: baseline;
`;

const ItemTable = styled.table`
　width: 100%;
　border-collapse: collapse;
`;

const Item = styled.tr`
  &:nth-child(odd) {
    background: #eef;
  }
`;

const Adjustment = styled(Item)`
  color: #009;
`;

const Total = styled.tr`
  background: #ccf;
  font-weight: bold;
`;

const ItemTitle = styled.td`
  width: 60%;
  padding: 0.5rem 0.7rem;
`;

const ItemAmount = styled.td`
  width: 15%;
  text-align: right;
  padding: 0.5rem 0.7rem;
`;

const DeleteColumn = styled.td`
  width: 10%;
  text-align: right;
  font-size: smaller;
  padding: 0.5rem 0.7rem;
`;

const DeleteButton = styled.button`
  background: none;
  border: none;
  color: #666;
  &:hover {
    color: #900;
  }
`;

function Month() {
    const { params: { year, month } } = useRouteMatch();
    useEffect(() => {
        fetchRecords();
    }, []);

    const [items, setItems] = useState([]);
    const [adjustment, setAdjustment] = useState({ amounts: [0, 0] });
    const [total, setTotal] = useState({ amounts: [0, 0] });

    function fetchRecords() {
        setItems([
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
        const ans = window.confirm('Delete the item?');
        if (ans) {
        } else {
        }
    }

    return (
        <div className="Month">
            <Header>
                <Title>{year}/{month}</Title>
                <NewForm />
            </Header>
            <ItemTable>
                <tbody>
                    {items.map((s, i) =>
                        <Item key={i}>
                            <ItemTitle>{s.title}</ItemTitle>
                            {s.amounts.map((a, i) =>
                                <ItemAmount key={i}>{number(a)}</ItemAmount>
                            )}
                            <DeleteColumn>
                                <DeleteButton onClick={deleteItem}>Delete</DeleteButton>
                            </DeleteColumn>
                        </Item>
                    )}
                    <Adjustment>
                        <ItemTitle>Adjustment</ItemTitle>
                        {adjustment.amounts.map((a, i) =>
                            <ItemAmount key={i}>{a}</ItemAmount>
                        )}
                        <td></td>
                    </Adjustment>
                    <Total>
                        <ItemTitle>Total</ItemTitle>
                        <ItemAmount>{total.amounts[0]}</ItemAmount>
                        <ItemAmount>{total.amounts[1]}</ItemAmount>
                        <td></td>
                    </Total>
                </tbody>
            </ItemTable>
        </div>
    );
}

export default Month;
