import React, { useState } from 'react';
import styled from 'styled-components';

const Container = styled.div`
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  align-content: flex-end;
  align-items: flex-end;
  position: relative;
`;

const ToggleButton = styled.button`
  display: block;
  width: 8rem;
`;

const Form = styled.form`
  background: #fff;
  border: 1pt #ccc solid;
  padding: 1rem;
  position: absolute;
  top: 2rem;
`;

const FormBody = styled.dl`
  display: flex;
  display: -webkit-flex;
  width: 20rem;
  flex-wrap: wrap;
  justify-content: flex-end;
`;

const InputLabel = styled.dt`
  width: 20%;
  border-bottom: solid #ccc 1pt;
  margin: 0 0 1rem 0;
  color: #666;
`;

const InputField = styled.dd`
  width: 80%;
  border-bottom: solid #ccc 1pt;
  margin: 0 0 1rem 0;        
`;

const Input = styled.input`
  border: none;
  width: 100%;
  outline: none;
`;

const SubmitButton = styled.button`
  display: block;
  margin: -0.5rem auto 0;
  width: 8rem;
`;

function NewForm({ reload }) {
  const [enabled, setEnabled] = useState();
  const [date, setDate] = useState(new Date().toISOString().slice(0, 10));

  async function handleSubmit(event) {
    event.preventDefault();
    const res = await fetch(`/v1/items`, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'ID-Token': sessionStorage.getItem('idToken'),
      },
      body: JSON.stringify({
        date: event.target.date.value,
        name: event.target.name.value,
        amount: parseInt(event.target.amount.value),
      }),
    });
    if (res.status === 204) {
      toggleEnabled();
      reload();
    } else {
      console.error("failed to post");
    }
  }

  function toggleEnabled() {
    setEnabled(!enabled);
  }

  function handleDateChange(event) {
    setDate(event.target.value);
  }

  return (
    <Container>
      <ToggleButton onClick={toggleEnabled}>New</ToggleButton>
      {enabled ? (
        <Form onSubmit={handleSubmit}>
          <FormBody>
            <InputLabel>Date</InputLabel>
            <InputField><Input name="date" type="date" value={date} onChange={handleDateChange}></Input></InputField>
            <InputLabel>Title</InputLabel>
            <InputField><Input name="name" type="text"></Input></InputField>
            <InputLabel>Amount</InputLabel>
            <InputField><Input name="amount"></Input></InputField>
          </FormBody>
          <SubmitButton type="submit">Submit</SubmitButton>
        </Form>
      ) : []}
    </Container>
  );
}

export default NewForm;
