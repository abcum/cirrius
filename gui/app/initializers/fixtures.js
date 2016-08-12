import Ember from "ember";
import FixtureAdapter from "gui/adapters/-fixture";
import FixtureSerializer from "gui/serializers/-fixture";
export default {
    name: 'fixtures',
    initialize: function(application) {
        application.register('adapter:-fixture', FixtureAdapter);
        application.register('serializer:-fixture', FixtureSerializer);
    }
};