import Ember from "ember";
import DS from "ember-data";

export default DS.FixtureSerializer = DS.JSONSerializer.extend({

    isNewSerializerAPI: true,

    serialize: function(snapshot, options={}) {

        let json = {};

        json.id = snapshot.id;

        snapshot.eachAttribute( (key, attribute) => {

            if (!attribute.options.readOnly) this.serializeAttribute(snapshot, json, key, attribute);

        });

        snapshot.eachRelationship( (key, relationship) => {

            if (relationship.kind === 'hasMany') {
                if (!relationship.options.readOnly) this.serializeHasMany(snapshot, json, relationship);
            }

            if (relationship.kind === 'belongsTo') {
                if (!relationship.options.readOnly) this.serializeBelongsTo(snapshot, json, relationship);
            }

        });

        return json;

    },

    normalizeSingleResponse: function(store, type, json, id, method) {

        if (json.meta && json.meta.links) {
            json.data.links = json.meta.links;
        }

        if (json.data) json.data.meta = json.meta || {};

        return this._super(store, type, json.data, id, method);

    },

    normalizeArrayResponse: function(store, type, json, id, method) {

        if (json.meta && json.meta.links) {
            for (var item of json.data) {
                item.links = json.meta.links;
            }
        }

        if (json.data) json.data.meta = json.meta || {};

        return this._super(store, type, json.data, id, method);

    },

});