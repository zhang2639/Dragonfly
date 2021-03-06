/*
 * Copyright The Dragonfly Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cdn

import (
	"context"
	"net/http"

	errorType "github.com/Dragonfly/pkg/errortypes"
	"github.com/Dragonfly/pkg/httputils"
	"github.com/Dragonfly/supernode/util"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// download downloads the file from the original address and
// sets the "Range" header to the undownloaded file range.
//
// If the returned error is nil, the Response will contain a non-nil
// Body which the caller is expected to close.
func (cm *Manager) download(ctx context.Context, taskID, url string, headers map[string]string,
	startPieceNum int, httpFileLength int64, pieceContSize int32) (*http.Response, error) {
	var checkCode = http.StatusOK

	if startPieceNum > 0 {
		breakRange, err := util.CalculateBreakRange(startPieceNum, int(pieceContSize), httpFileLength)
		if err != nil {
			return nil, errors.Wrapf(errorType.ErrInvalidValue, "failed to calculate the breakRange: %v", err)
		}

		if headers == nil {
			headers = make(map[string]string)
		}
		headers["Range"] = httputils.ConstructRangeStr(breakRange)
		checkCode = http.StatusPartialContent
	}

	logrus.Infof("start to download for taskId(%s) with fileUrl: %s header: %v checkCode: %d", taskID, url, headers, checkCode)
	return cm.originClient.Download(url, headers, checkCode)
}
