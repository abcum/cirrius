#!/usr/bin/env sh

folder="$(pwd)"

versions=10

packages=()

bow=()
bow+=('async:::dist/async.min.js')
bow+=('bluebird:::js/browser/bluebird.min.js')
bow+=('lodash:::dist/lodash.min.js')
bow+=('underscore:::underscore-min.js')

npm=()
npm+=('accounting:::accounting.min.js')
npm+=('chance:::dist/chance.min.js')
npm+=('deepmerge:::index.js')
npm+=('hashids:::dist/hashids.min.js')
npm+=('immutable:::dist/immutable.min.js')
npm+=('jsondiffpatch:::public/build/jsondiffpatch-full.min.js')
npm+=('merge:::merge.min.js')
npm+=('moment:::min/moment.min.js')
npm+=('odiff:::dist/odiff.umd.js')
npm+=('promiz:::promiz.min.js')
npm+=('q:::q.js')
npm+=('sugar:::dist/sugar.min.js')
npm+=('validator:::validator.min.js')

for index in "${bow[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then

            echo "--- $name"#"$version"

            # bower cache clean

			cd $folder
			bower install $name"#"$version
			cd "$folder/bower_components/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js" || { echo 'File not found'; exit 1; }
			rm -rf "$folder/bower_components/$name"

		fi

	done

done

for index in "${npm[@]}" ; do

	name="${index%%:::*}"
	path="${index##*:::}"

    npm view $name versions --json | grep '"' | cut -d '"' -f2 | tail -$versions | while read version; do

    	if [ ! -f "$folder/npm/modules/$name/$version.js" ]; then

            echo "--- $name"@"$version"

            # npm cache clean

			cd $folder
			npm install $name"@"$version
			cd "$folder/node_modules/$name"
			cp "$path" "$folder/npm/modules/$name/$version.js" || { echo 'File not found'; exit 1; }
			rm -rf "$folder/node_modules/$name"

		fi

	done

done

# rm -rf "$folder/bower_components/"
# rm -rf "$folder/node_modules/"

exit 0
