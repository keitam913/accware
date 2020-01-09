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

function NewForm() {
        const [enabled, setEnabled] = useState();

        function handleSubmit() {
                console.log('submit');
        }

        function toggleEnabled() {
                setEnabled(!enabled);
        }

        return (
                <Container>
                        <ToggleButton onClick={toggleEnabled}>New</ToggleButton>
                        {enabled ? (
                                <Form onSubmit={handleSubmit}>
                                        <FormBody>
                                                <InputLabel>Title</InputLabel>
                                                <InputField><Input name="inputTitle"></Input></InputField>
                                                <InputLabel>Amount</InputLabel>
                                                <InputField><Input name="inputAmount"></Input></InputField>
                                        </FormBody>
                                        <SubmitButton type="submit">Submit</SubmitButton>
                                </Form>
                        ) : []}
                </Container>
        );
}

export default NewForm;
