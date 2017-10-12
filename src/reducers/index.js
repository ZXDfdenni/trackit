import { combineReducers } from 'redux';
import aws from './aws';
import gcp from './gcp';

export default combineReducers({
  aws,
  gcp
});
