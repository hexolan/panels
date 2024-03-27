# Copyright 2023 Declan Teevan
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from typing import Type, List

from comment_service.models.service import CommentRepository, Comment, CommentCreate, CommentUpdate


class ServiceRepository(CommentRepository):
    """The comment service repository.

    Attributes:
        _repo (Type[CommentRepository]): The downstream repository (Redis Repository -> DB Repository).
    
    """
    def __init__(self, downstream_repo: Type[CommentRepository]) -> None:
        self._repo = downstream_repo

    async def get_comment(self, comment_id: int) -> Comment:
        return await self._repo.get_comment(comment_id)

    async def get_post_comments(self, post_id: str) -> List[Comment]:
        # todo: pagination
        return await self._repo.get_post_comments(post_id)
    
    async def create_comment(self, data: CommentCreate) -> Comment:
        return await self._repo.create_comment(data)
    
    async def update_comment(self, comment_id: int, data: CommentUpdate) -> Comment:
        return await self._repo.update_comment(comment_id, data)
    
    async def delete_comment(self, comment_id: int) -> None:
        await self._repo.delete_comment(comment_id)