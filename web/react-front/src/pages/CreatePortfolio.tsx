import React from 'react';

const CreatePortfolio = () => {
    return (
        <form onSubmit={submit}>
        <h1 className="h3 mb-3 fw-normal">Choose portfolio parameters</h1>

        <input className="form-control" placeholder="Name" required
               onChange = {e => setName(e.target.value)} 
        />

        <input className="form-control" placeholder="Description" required
               onChange = {e => setDescription(e.target.value)}
        />
        <label>
            Public
        <input type="checkbox" className="form-control" placeholder="Public" required
               onChange = {e => setEmail(e.target.value)}
        />
        </label>
        <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
    );
};

export default CreatePortfolio;