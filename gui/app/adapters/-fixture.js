import Ember from "ember";
import DS from "ember-data";

export default DS.FixtureAdapter = DS.Adapter.extend({

    defaultSerializer: '-fixture',

    shouldReloadAll(store, snapshots) {
        return false;
    },

    shouldReloadRecord(store, snapshot) {
        return false;
    },

    shouldBackgroundReloadAll(store, snapshots) {
        return true;
    },

    shouldBackgroundReloadRecord(store, snapshot) {
        return true;
    },

    generateIdForRecord(store, type) {

        return type + ':' + Math.random().toString(16).slice(2);

    },

    findAll(store, type, since) {

        let data = this.fixtures(type);

        return this.requests(data, { total: data.length });

    },

    findMany(store, type, ids, snapshots) {

        let data = this.fixtures(type).filter(function(item) {
            return ids.indexOf(item.id) !== -1;
        });

        return this.requests(data, { total: data.length });

    },

    findRecord(store, type, id, snapshot) {

        let data = this.fixtures(type).findBy('id', id);

        return this.requests(data, { total: data.length });

    },

    query(store, type, query) {

    	if (query.filter) {

    		let data = this.fixtures(type);

	        let part = data.filter(item => {
	        	return Object.keys(query.filter).every(key => {
	        		return item[key] === query.filter[key];
	        	});
	        });

	        return this.requests(part, { total: data.length });

    	} else {

	        query.start = Math.max(0, query.start-1);

	        let data = this.fixtures(type);

	        let part = data.slice(query.start, query.start + query.limit);

	        return this.requests(part, { total: data.length });

	    }

    },

    createRecord(store, type, snapshot) {

        let record = store.serializerFor(type.modelName).serialize(snapshot);

        Ember.A(type.FIXTURES).removeObject( Ember.A(type.FIXTURES).findBy('id', record.id) );

        Ember.A(type.FIXTURES).pushObject(record);

        return this.requests(record);

    },

    updateRecord(store, type, snapshot) {

        let record = store.serializerFor(type.modelName).serialize(snapshot);

        Ember.A(type.FIXTURES).removeObject( Ember.A(type.FIXTURES).findBy('id', record.id) );

        Ember.A(type.FIXTURES).pushObject(record);

        return this.requests(record);

    },

    deleteRecord(store, type, snapshot) {

        let record = store.serializerFor(type.modelName).serialize(snapshot);

        Ember.A(type.FIXTURES).removeObject( Ember.A(type.FIXTURES).findBy('id', record.id) );

        return this.requests(null);

    },

    fixtures(type) {

        return Ember.copy(type.FIXTURES, true);

    },

    requests(data, meta) {

        // Return a promise and resolve it
        // to simulate remote promise

        return new Ember.RSVP.Promise( (resolve, reject) => {

            resolve({ meta:meta, data:data });

        });

    },

});
