import { put, call, all } from 'redux-saga/effects';
import { getAccountsSaga, newAccountSaga, newExternalSaga } from '../accountsSaga';
import { getToken } from '../../misc';
import API from '../../../api';
import Constants from '../../../constants';

const token = "42";

describe("Accounts Saga", () => {

  describe("Get Accounts", () => {

    const accounts = ["account1", "account2"];
    const validResponse = { success: true, data: accounts };
    const invalidResponse = { success: true, accounts };
    const noResponse = { success: false };

    it("handless saga with valid data", () => {

      let saga = getAccountsSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.getAccounts, token));

      expect(saga.next(validResponse).value)
        .toEqual(all([
          put({ type: Constants.AWS_GET_ACCOUNTS_SUCCESS, accounts })
        ]));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with invalid data", () => {

      let saga = getAccountsSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.getAccounts, token));

      expect(saga.next(invalidResponse).value)
        .toEqual(put({ type: Constants.AWS_GET_ACCOUNTS_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with no response", () => {

      let saga = getAccountsSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.getAccounts, token));

      expect(saga.next(noResponse).value)
        .toEqual(put({ type: Constants.AWS_GET_ACCOUNTS_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

  });

  describe("New Account", () => {

    const account = {roleArn: "roleArn"};
    const validResponse = { success: true, data: account };
    const invalidResponse = { success: true, account };
    const noResponse = { success: false };

    it("handless saga with valid data", () => {

      let saga = newAccountSaga({account});

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newAccount, account, token));

      expect(saga.next(validResponse).value)
        .toEqual(all([
          put({ type: Constants.AWS_NEW_ACCOUNT_SUCCESS, account }),
          put({ type: Constants.AWS_GET_ACCOUNTS })
        ]));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with invalid data", () => {

      let saga = newAccountSaga({account});

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newAccount, account, token));

      expect(saga.next(invalidResponse).value)
        .toEqual(put({ type: Constants.AWS_NEW_ACCOUNT_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with no response", () => {

      let saga = newAccountSaga({account});

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newAccount, account, token));

      expect(saga.next(noResponse).value)
        .toEqual(put({ type: Constants.AWS_NEW_ACCOUNT_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

  });

  describe("New External", () => {

    const external = "external";
    const validResponse = { success: true, data: { external } };
    const invalidResponse = { success: true, external };
    const noResponse = { success: false };

    it("handless saga with valid data", () => {

      let saga = newExternalSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newExternal, token));

      expect(saga.next(validResponse).value)
        .toEqual(all([
          put({ type: Constants.AWS_NEW_EXTERNAL_SUCCESS, external })
        ]));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with invalid data", () => {

      let saga = newExternalSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newExternal, token));

      expect(saga.next(invalidResponse).value)
        .toEqual(put({ type: Constants.AWS_NEW_EXTERNAL_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

    it("handless saga with no response", () => {

      let saga = newExternalSaga();

      expect(saga.next().value)
        .toEqual(getToken());

      expect(saga.next(token).value)
        .toEqual(call(API.AWS.Accounts.newExternal, token));

      expect(saga.next(noResponse).value)
        .toEqual(put({ type: Constants.AWS_NEW_EXTERNAL_ERROR, error: Error("Error with request") }));

      expect(saga.next().done).toBe(true);

    });

  });

});
