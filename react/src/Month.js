import React, { useState, useEffect } from 'react';
import { useRouteMatch, Redirect } from 'react-router-dom';
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

const Delete = styled.button`
  background: none;
  border: none;
  color: #666;
  &:hover {
    color: #900;
  }
`;

function DeleteButton({ itemId, reload }) {
  async function deleteItem() {
    const ans = window.confirm('Delete the item?');
    if (!ans) {
      return;
    }
    await fetch(`/v1/items/${itemId}`, {
      method: 'DELETE',
      mode: 'cors',
      headers: {
        'ID-Token': sessionStorage.getItem('idToken'),
      },
    });
    reload();
  }
  return <Delete onClick={deleteItem}>Delete</Delete>
}

function Month() {
  const { params: { year, month } } = useRouteMatch();
  useEffect(() => {
    updateRecords();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const [items, setItems] = useState([]);
  const [adjustment, setAdjustment] = useState({ amounts: [0, 0] });
  const [total, setTotal] = useState({ amounts: [0, 0] });
  const [authenticated, setAuthenticated] = useState(true);

  async function updateRecords() {
    const res = await fetch(`/v1/accounts/${year}/${month}`, {
      method: 'GET',
      headers: {
        'ID-Token': sessionStorage.getItem('idToken'),
      },
    });

    if (res.status === 401) {
      setAuthenticated(false);
      return
    }

    const j = await res.json()
    const pids = Object.keys(j.persons);

    const ni = [];
    j.items.forEach((item) => {
      let am = [0, 0];
      for (let i = 0; i < pids.length; i++) {
        if (pids[i] === item.personId) {
          am[i] = item.amount;
        }
      }
      ni.push({
        id: item.id,
        title: item.name,
        amounts: am,
      })
    });
    setItems(ni);

    const na = [];
    j.adjustments.forEach((a) => {
      for (let i = 0; i < pids.length; i++) {
        if (pids[i] === a.personId) {
          na[i] = a.amount;
        }
      }
    });
    setAdjustment({
      amounts: na,
    });

    const nt = [];
    j.totals.forEach((a) => {
      for (let i = 0; i < pids.length; i++) {
        if (pids[i] === a.personId) {
          nt[i] = a.amount;
        }
      }
    });
    setTotal({
      amounts: nt,
    });
  }

  function number(n) {
    if (n === 0) {
      return "";
    } else {
      return n.toString();
    }
  }

  if (!authenticated) {
    return <Redirect to="/login" />;
  }

  return (
    <div className="Month">
      <Header>
        <Title>{year}/{month}</Title>
        <NewForm reload={updateRecords} />
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
                <DeleteButton reload={updateRecords} itemId={s.id}>Delete</DeleteButton>
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
