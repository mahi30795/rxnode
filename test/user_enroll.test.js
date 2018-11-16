/**
 * Testing Script to test user enrollment
 *
 *  @author Sachu Shaji Abraham <sachu.shaji@netobjex.com>
 */

/* global describe,it,before */


require('mocha');
const { expect } = require('chai');

const axios = require('axios');

let Enrollment1; // variable to store response of user one
let Enrollment2; // variable to store response of user two
let Enrollment3; // variable to store response of user three

before('Running pre configurations', async function enroll() {
  this.timeout(0);
  // Running Enrollment One
  await axios({
    method: 'post',
    url: ' http://localhost:4000/users',
    data: {
      username: 'Adam',
      orgName: 'Org1',
    },
  }).then((res) => {
    Enrollment1 = res;
  }).catch((err) => {
    Enrollment1 = err;
  });
  // Running Enrollment Two
  await axios({
    method: 'post',
    url: ' http://localhost:4000/users',
    data: {
      username: 'Max',
      orgName: 'Org2',
    },
  }).then((res) => {
    Enrollment2 = res;
  }).catch((err) => {
    Enrollment2 = err;
  });
  // Running Enrollment Three
  await axios({
    method: 'post',
    url: ' http://localhost:4000/users',
    data: {
      username: 'Kate',
      orgName: 'Org3',
    },
  }).then((res) => {
    Enrollment3 = res;
  }).catch((err) => {
    Enrollment3 = err;
  });
});

describe('Testing the enrollment of the users', () => {
  it('it should succesfully register the users one ', () => {
    expect(Enrollment1)
      .to.have.property('status')
      .equals(200);
  });
  it('it should succesfully register the users two', () => {
    expect(Enrollment2)
      .to.have.property('status')
      .equals(200);
  });
  it('it should succesfully register the users three', () => {
    expect(Enrollment3)
      .to.have.property('status')
      .equals(200);
  });
});
