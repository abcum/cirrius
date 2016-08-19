import Ember from "ember";

export default Ember.Controller.extend({

    file: Ember.inject.controller(),
    request: Ember.inject.controller(),

    actions: {

        create: function() {

            this.store
                .createRecord('file', { project: this.get('model') })
                .save()
                .then(model => {
                    this.transitionToRoute('file', model);
                })
            ;

        },

    }

});
